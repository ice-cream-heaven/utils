// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePaySupplierProduct = "pay_supplier_product"

// PaySupplierProduct mapped from table <pay_supplier_product>
type PaySupplierProduct struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`           // ID
	SupplierID int64                 `gorm:"column:supplier_id;not null;comment:供应商 id" json:"supplier_id"`          // 供应商 id
	ProductID  int64                 `gorm:"column:product_id;not null;comment:产品 id" json:"product_id"`             // 产品 id
	Status     int64                 `gorm:"column:status;not null;default:1;comment:状态:1 启用，2 关闭" json:"status"`    // 状态:1 启用，2 关闭
	CreateTime int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"` // 创建时间
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"` // 更新时间
	Delete_    soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`              // 删除状态:0 未删除，1 已删除
}

// TableName PaySupplierProduct's table name
func (*PaySupplierProduct) TableName() string {
	return TableNamePaySupplierProduct
}
