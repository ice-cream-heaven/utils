package db

import (
	"fmt"
	"github.com/ice-cream-heaven/log"
	"github.com/ice-cream-heaven/utils/anyx"
	"github.com/ice-cream-heaven/utils/candy"
	"github.com/ice-cream-heaven/utils/json"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"time"
)

type ModelScoop[M any] struct {
	Scoop

	m M

	rt     reflect.Type
	fields []*Field
}

func NewModelScoop[M any](db *gorm.DB) *ModelScoop[M] {
	scoop := &ModelScoop[M]{
		Scoop: Scoop{
			_db: db.Session(&gorm.Session{}),
		},
	}

	scoop.inc()

	return scoop
}

// ——————————条件——————————

func (p *ModelScoop[M]) Select(fields ...string) *ModelScoop[M] {
	p.selects = append(p.selects, fields...)
	return p
}

func (p *ModelScoop[M]) Where(args ...interface{}) *ModelScoop[M] {
	p.cond.Where(args...)
	return p
}

func (p *ModelScoop[M]) Equal(column string, value interface{}) *ModelScoop[M] {
	p.cond.where(column, value)
	return p
}

func (p *ModelScoop[M]) NotEqual(column string, value interface{}) *ModelScoop[M] {
	p.cond.where(column, "!= ", value)
	return p
}

func (p *ModelScoop[M]) In(column string, values interface{}) *ModelScoop[M] {
	vo := EnsureIsSliceOrArray(values)
	if vo.Len() == 0 {
		p.cond.where(false)
		return p
	}
	p.cond.where(column, "IN", UniqueSlice(vo.Interface()))
	return p
}

func (p *ModelScoop[M]) NotIn(column string, values interface{}) *ModelScoop[M] {
	vo := EnsureIsSliceOrArray(values)
	if vo.Len() == 0 {
		return p
	}
	p.cond.where(column, "NOT IN", UniqueSlice(vo.Interface()))
	return p
}

func (p *ModelScoop[M]) Like(column string, value string) *ModelScoop[M] {
	p.cond.where(column, "LIKE", "%"+value+"%")
	return p
}

func (p *ModelScoop[M]) LeftLike(column string, value string) *ModelScoop[M] {
	p.cond.where(column, "LIKE", value+"%")
	return p
}

func (p *ModelScoop[M]) RightLike(column string, value string) *ModelScoop[M] {
	p.cond.where(column, "LIKE", "%"+value)
	return p
}

func (p *ModelScoop[M]) NotLike(column string, value string) *ModelScoop[M] {
	p.cond.where(column, "NOT LIKE", "%"+value+"%")
	return p
}

func (p *ModelScoop[M]) NotLeftLike(column string, value string) *ModelScoop[M] {
	p.cond.where(column, "NOT LIKE", value+"%")
	return p
}

func (p *ModelScoop[M]) NotRightLike(column string, value string) *ModelScoop[M] {
	p.cond.where(column, "NOT LIKE", "%"+value)
	return p
}

func (p *ModelScoop[M]) Between(column string, min, max interface{}) *ModelScoop[M] {
	p.cond.whereRaw(fmt.Sprintf(quoteFieldName(column))+" BETWEEN ? AND ?", min, max)
	return p
}

func (p *ModelScoop[M]) NotBetween(column string, min, max interface{}) *ModelScoop[M] {
	p.cond.whereRaw(fmt.Sprintf(quoteFieldName(column))+" NOT BETWEEN ? AND ?", min, max)
	return p
}

func (p *ModelScoop[M]) Unscoped(b ...bool) *ModelScoop[M] {
	if len(b) == 0 {
		p.unscoped = true
		return p
	}
	p.unscoped = b[0]
	return p
}

func (p *ModelScoop[M]) Limit(limit uint64) *ModelScoop[M] {
	p.limit = limit
	return p
}

func (p *ModelScoop[M]) Offset(offset uint64) *ModelScoop[M] {
	p.offset = offset
	return p
}

func (p *ModelScoop[M]) Group(fields ...string) *ModelScoop[M] {
	p.groups = append(p.groups, fields...)
	return p
}

func (p *ModelScoop[M]) Order(fields ...string) *ModelScoop[M] {
	p.orders = append(p.orders, fields...)
	return p
}

func (p *ModelScoop[M]) Ignore(b ...bool) *ModelScoop[M] {
	if len(b) == 0 {
		p.ignore = true
		return p
	}

	p.ignore = b[0]

	return p
}

func (p *ModelScoop[M]) Replace(b ...bool) *ModelScoop[M] {
	if len(b) == 0 {
		p.replace = true
		return p
	}

	p.replace = b[0]

	return p
}

// ——————————操作——————————

func (p *ModelScoop[M]) First() (*M, error) {
	p.inc()
	defer p.dec()

	var m M
	err := p.Scoop.First(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (p *ModelScoop[M]) Find() ([]*M, error) {
	p.inc()
	defer p.dec()

	var ms []*M
	err := p.Scoop.Find(&ms).Error
	if err != nil {
		return nil, err
	}

	return ms, nil
}

func (p *ModelScoop[M]) Create(m *M) error {
	p.inc()
	defer p.dec()

	return p.Scoop.Create(m).Error
}

type FirstOrCreateResult[M any] struct {
	IsCreated bool
	Error     error

	Object *M
}

func (p *ModelScoop[M]) FirstOrCreate(m *M) *FirstOrCreateResult[M] {
	p.inc()
	defer p.dec()

	var mm M
	err := p.Scoop.First(&mm).Error
	if err != nil {
		if p.IsNotFound(err) {
			err = p.Scoop.Create(m).Error
			if err != nil {
				return &FirstOrCreateResult[M]{
					Error: err,
				}
			}
			return &FirstOrCreateResult[M]{
				IsCreated: true,
				Object:    m,
			}
		}
		return &FirstOrCreateResult[M]{
			Error: err,
		}
	}
	return &FirstOrCreateResult[M]{
		Object: &mm,
	}
}

type UpdateOrCreateResult[M any] struct {
	IsCreated bool
	Error     error

	Object *M
}

func (p *ModelScoop[M]) UpdateOrCreate(values map[string]interface{}, m *M) *UpdateOrCreateResult[M] {
	p.inc()
	defer p.dec()

	var mm M
	err := p.Scoop.First(&mm).Error
	if err != nil {
		if p.IsNotFound(err) {
			err = p.Scoop.Create(m).Error
			if err != nil {
				return &UpdateOrCreateResult[M]{
					Error: err,
				}
			}
			return &UpdateOrCreateResult[M]{
				IsCreated: true,
				Object:    m,
			}
		}

		return &UpdateOrCreateResult[M]{
			Error: err,
		}
	}

	err = p.Scoop.update(values).Error
	if err != nil {
		return &UpdateOrCreateResult[M]{
			Error: err,
		}
	}

	err = p.Scoop.First(&mm).Error
	if err != nil {
		return &UpdateOrCreateResult[M]{
			Error: err,
		}
	}

	anyx.DeepCopy(&mm, m)

	return &UpdateOrCreateResult[M]{
		Object: &mm,
	}
}

type CreateNotExistResult[M any] struct {
	IsCreated bool
	Error     error

	Object *M
}

func (p *ModelScoop[M]) CreateNotExist(m *M) *CreateNotExistResult[M] {
	p.inc()
	defer p.dec()

	var mm M
	err := p.Scoop.First(&mm).Error
	if err != nil {
		if p.IsNotFound(err) {
			err = p.Scoop.Create(m).Error
			if err != nil {
				return &CreateNotExistResult[M]{
					Error: err,
				}
			}
			return &CreateNotExistResult[M]{
				IsCreated: true,
				Object:    m,
			}
		}
		return &CreateNotExistResult[M]{
			Error: err,
		}
	}

	anyx.DeepCopy(&mm, m)

	return &CreateNotExistResult[M]{
		Object: &mm,
	}
}

func (p *ModelScoop[M]) Chunk(size uint64, fc func(tx *Scoop, out []*M, offset uint64) error) *ChunkResult {
	p.inc()
	defer p.dec()

	var out []*M
	return p.Scoop.Chunk(&out, size, func(tx *Scoop, offset uint64) error {
		return fc(tx, out, offset)
	})
}

type CreateOrUpdateResult[M any] struct {
	Error  error
	Object *M

	Created bool
	Updated bool
}

func (p *ModelScoop[M]) CreateOrUpdate(values map[string]interface{}, m *M) *CreateOrUpdateResult[M] {
	p.inc()
	defer p.dec()

	var old M
	err := p.Scoop.First(&old).Error
	if err != nil {
		if !p.IsNotFound(err) {
			log.Errorf("err:%s", err)
			return &CreateOrUpdateResult[M]{
				Error: err,
			}
		}
		// 创建

		err = p.Scoop.Create(m).Error
		if err != nil {
			log.Errorf("err:%s", err)
			return &CreateOrUpdateResult[M]{
				Error: err,
			}
		}

		return &CreateOrUpdateResult[M]{
			Object:  m,
			Created: true,
		}
	} else if len(values) > 0 {
		// 更新
		err = p.Scoop.Updates(values).Error
		if err != nil {
			log.Errorf("err:%s", err)
			return &CreateOrUpdateResult[M]{
				Error: err,
			}
		}

		err = p.Scoop.First(&old).Error
		if err != nil {
			if !p.Scoop.IsNotFound(err) {
				log.Errorf("err:%s", err)
			}
			return &CreateOrUpdateResult[M]{
				Error: err,
			}
		}

		anyx.DeepCopy(&old, m)

		return &CreateOrUpdateResult[M]{
			Object:  &old,
			Updated: true,
		}
	} else {

		anyx.DeepCopy(&old, m)

		return &CreateOrUpdateResult[M]{
			Object: &old,
		}
	}
}

func (p *ModelScoop[M]) CreateInBatches(value []*M, batchSize int) *CreateInBatchesResult {
	p.inc()
	defer p.dec()

	result := &CreateInBatchesResult{}

	do := func(values []*M) error {
		session := p._db.Session(&gorm.Session{})

		sqlRaw := log.GetBuffer()
		defer log.PutBuffer(sqlRaw)

		if p.replace {
			sqlRaw.WriteString("REPLACE ")
		} else {
			sqlRaw.WriteString("INSERT ")
		}

		if p.ignore {
			sqlRaw.WriteString("IGNORE ")
		}

		sqlRaw.WriteString("INTO ")
		sqlRaw.WriteString(p.table)

		sqlRaw.WriteString(" (")
		for _, field := range p.fields {
			sqlRaw.WriteString(field.GormField)
			sqlRaw.WriteString(",")
		}
		sqlRaw.Truncate(sqlRaw.Len() - 1)
		sqlRaw.WriteString(") VALUES ")

		for _, v := range values {
			vv := reflect.ValueOf(v)
			if vv.Kind() == reflect.Ptr {
				vv = vv.Elem()
			}

			sqlRaw.WriteString("(")
			for _, field := range p.fields {
				switch field.GormField {
				case "id":
					if vv.FieldByName(field.Name).IsZero() {
						sqlRaw.WriteString("NULL")
						sqlRaw.WriteString(",")
						continue
					}

				case "created_at", "updated_at":
					sqlRaw.WriteString("NOW()")
					sqlRaw.WriteString(",")

				default:
					// do nothing
				}

				if vv.FieldByName(field.Name).IsZero() {
					switch field.Type {
					case reflect.String:
						sqlRaw.WriteString("''")

					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						sqlRaw.WriteString("0")

					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						sqlRaw.WriteString("0")

					case reflect.Float32, reflect.Float64:
						sqlRaw.WriteString("0")

					case reflect.Bool:
						sqlRaw.WriteString("0")

					case reflect.Ptr, reflect.Struct:
						sqlRaw.WriteString("{}")

					case reflect.Slice, reflect.Array:
						sqlRaw.WriteString("[]")

					default:
						sqlRaw.WriteString("NULL")

					}
					sqlRaw.WriteString(",")
					continue
				}

				switch field.Type {
				case reflect.String:
					sqlRaw.WriteString("'")
					sqlRaw.WriteString(vv.FieldByName(field.Name).String())
					sqlRaw.WriteString("'")

				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					sqlRaw.WriteString(strconv.FormatInt(vv.FieldByName(field.Name).Int(), 10))

				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					sqlRaw.WriteString(strconv.FormatUint(vv.FieldByName(field.Name).Uint(), 10))

				case reflect.Float32, reflect.Float64:
					sqlRaw.WriteString(strconv.FormatFloat(vv.FieldByName(field.Name).Float(), 'f', -1, 64))

				case reflect.Bool:
					if vv.FieldByName(field.Name).Bool() {
						sqlRaw.WriteString("1")
					} else {
						sqlRaw.WriteString("0")
					}

				default:
					sqlRaw.WriteString("'")
					sqlRaw.Write(json.MustMarshal(vv.FieldByName(field.Name).Interface()))
					sqlRaw.WriteString("'")
				}

				sqlRaw.WriteString(",")
			}

			sqlRaw.WriteString("),")
		}

		sqlRaw.Grow(1)

		start := time.Now()
		session.Exec(sqlRaw.String())

		getDefaultLogger().Log(p.depth, start, func() (sql string, rowsAffected int64) {
			return sqlRaw.String(), session.RowsAffected
		}, session.Error)

		result.RowsAffected += session.RowsAffected

		if p.ignore {
			// do nothing
		} else if session.Error != nil {
			result.Error = session.Error
			return session.Error
		}

		return nil
	}

	for _, values := range candy.Chunk(value, batchSize) {
		err := do(values)
		if err != nil {
			return result
		}
	}

	return result
}
