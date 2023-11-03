package db_test

import (
	"github.com/ice-cream-heaven/utils/db"
	"testing"
)

func TestCond(t *testing.T) {
	// t.Log(cond.Clean().Where("a", 1).ToString())

	t.Log(db.OrWhere(map[string]any{
		"a": 1,
	}, map[string]any{
		"a": 2,
	}, map[string]any{
		"a": 3,
	}).ToString())

}

// (a = 1 and b =2 ) or (a = 2 and b = 1)
func TestSubCond(t *testing.T) {
	t.Log(db.OrWhere(db.Where(map[string]any{
		"a": 1,
		"b": 2,
	}), db.Where(map[string]any{
		"a": 2,
		"b": 3,
	})).ToString())
}

func TestLike(t *testing.T) {
	t.Log(db.Where("name", "like", "%a%").ToString())
}
