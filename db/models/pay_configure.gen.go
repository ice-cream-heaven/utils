// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayConfigure = "pay_configure"

// PayConfigure mapped from table <pay_configure>
type PayConfigure struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Code       string                `gorm:"column:code;not null;comment:配置编码" json:"code"`                // 配置编码
	Name       string                `gorm:"column:name;not null;comment:配置名称" json:"name"`                // 配置名称
	Value      string                `gorm:"column:value" json:"value"`
	CreateTime int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"` // 创建时间
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"` // 更新时间
	Default    int64                 `gorm:"column:_default;not null;comment:缺省状态:0 无，1 缺省｜不可删" json:"_default"`     // 缺省状态:0 无，1 缺省｜不可删
	Delete_    soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`              // 删除状态:0 未删除，1 已删除
}

// TableName PayConfigure's table name
func (*PayConfigure) TableName() string {
	return TableNamePayConfigure
}
