package db

import (
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type Tabler interface {
	TableName() string
}

type Model[M any] struct {
	db *Client
	m  M

	notFoundError error

	hasDeletedAt bool
	hasId        bool
	table        string
	rt           reflect.Type

	fields []*Field
}

func NewModel[M any](db *Client) *Model[M] {
	p := &Model[M]{
		db:            db,
		notFoundError: gorm.ErrRecordNotFound,

		rt: reflect.TypeOf(new(M)),
	}

	for p.rt.Kind() == reflect.Ptr {
		p.rt = p.rt.Elem()
	}

	p.hasId = hasId(p.rt)
	p.hasDeletedAt = hasDeleted(p.rt)
	p.table = getTableName(p.rt)

	for i := 0; i < p.rt.NumField(); i++ {
		ff := p.rt.Field(i)

		field := &Field{
			Name: ff.Name,
			Type: ff.Type.Kind(),
		}

		if tags, ok := ff.Tag.Lookup("gorm"); ok {
			for _, tag := range strings.Split(tags, ";") {
				if strings.HasPrefix(tag, "column") {
					field.GormField = strings.Split(tag, ":")[1]
				}
			}
		}

		if field.GormField == "" {
			field.GormField = Camel2UnderScore(field.Name)
		}

		p.fields = append(p.fields, field)
	}

	return p
}

func (p *Model[M]) SetNotFound(err error) *Model[M] {
	p.notFoundError = err

	return p
}

func (p *Model[M]) IsNotFound(err error) bool {
	return err == p.notFoundError || gorm.ErrRecordNotFound == err
}

func (p *Model[M]) NewScoop(tx ...*Scoop) *ModelScoop[M] {
	var db *gorm.DB
	if len(tx) == 0 || tx[0] == nil {
		db = p.db.db
	} else {
		db = tx[0]._db
	}

	scoop := NewModelScoop[M](db)
	scoop.hasDeletedAt = p.hasDeletedAt
	scoop.hasId = p.hasId
	scoop.table = p.table
	scoop.notFoundError = p.notFoundError
	scoop.rt = p.rt
	scoop.fields = p.fields

	return scoop
}

func (p *Model[M]) TableName() string {
	return p.table
}
