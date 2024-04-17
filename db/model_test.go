package db_test

import (
	"github.com/ice-cream-heaven/utils/db"
	"github.com/ice-cream-heaven/utils/db/models"
	"reflect"
	"testing"
)

func TestField(t *testing.T) {
	vv := reflect.TypeOf(&db.Field{})
	for vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	vvv := reflect.ValueOf(&db.Field{
		Name: "name",
	})
	for vvv.Kind() == reflect.Ptr {
		vvv = vvv.Elem()
	}

	for i := 0; i < vv.NumField(); i++ {
		t.Log(vv.Field(i).Name)
	}

	t.Log(vvv.FieldByNameFunc(func(s string) bool {
		return s == "Name"
	}))

	t.Log(vvv.FieldByName("Name").IsValid())
}

func TestMerge(t *testing.T) {
	d, err := db.New(&db.Config{
		Type:     "mysql",
		Address:  "127.0.0.1",
		Port:     3306,
		Name:     "xxpay_daily",
		Username: "root",
		Password: "HNEzz4fang.",
		Logger:   nil,
	})
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	err = d.AutoMigrate(&models.PayProductQueueConfig{})

	const (
		Id   = "id"
		Name = "name"
	)

	d.NewScoop().
		Where(Id, 1).
		Where(map[string]any{
			Id:   1,
			Name: "name",
		}).
		Equal(Id, 1).
		Updates(map[string]any{
			Id:   2,
			Name: "name33",
		})

	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
}
