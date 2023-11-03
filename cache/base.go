package cache

import "github.com/bytedance/sonic"

type baseCache struct {
	BaseCache
}

func (p *baseCache) GetJson(key string, j interface{}) error {
	value, err := p.Get(key)
	if err != nil {
		return err
	}

	return sonic.UnmarshalString(value, j)
}

func (p *baseCache) HGetJson(key, field string, j interface{}) error {
	value, err := p.HGet(key, field)
	if err != nil {
		return err
	}

	return sonic.UnmarshalString(value, j)
}

func newBaseCache(c BaseCache) Cache {
	return &baseCache{
		BaseCache: c,
	}
}
