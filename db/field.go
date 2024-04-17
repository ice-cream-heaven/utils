package db

import (
	"reflect"
)

type Field struct {
	Name      string
	Type      reflect.Kind
	GormField string
}
