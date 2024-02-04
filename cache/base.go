package cache

import (
	"github.com/ice-cream-heaven/utils/json"
	"time"
)

type baseCache struct {
	BaseCache
}

func (p *baseCache) Limit(key string, limit int64, timeout time.Duration) (bool, error) {
	cnt, err := p.Incr(key)
	if err != nil {
		return false, err
	}

	if cnt == 1 {
		_, err = p.Expire(key, timeout)
		if err != nil {
			return false, err
		}
	}

	if cnt > limit {
		return false, nil
	}

	return true, nil
}

func (p *baseCache) GetJson(key string, j interface{}) error {
	value, err := p.Get(key)
	if err != nil {
		return err
	}

	return json.UnmarshalString(value, j)
}

func (p *baseCache) HGetJson(key, field string, j interface{}) error {
	value, err := p.HGet(key, field)
	if err != nil {
		return err
	}

	return json.UnmarshalString(value, j)
}

func newBaseCache(c BaseCache) Cache {
	return &baseCache{
		BaseCache: c,
	}
}
