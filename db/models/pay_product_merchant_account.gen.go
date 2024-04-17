// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayProductMerchantAccount = "pay_product_merchant_account"

// PayProductMerchantAccount mapped from table <pay_product_merchant_account>
type PayProductMerchantAccount struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`           // ID
	MerchantID int64                 `gorm:"column:merchant_id;not null;comment:商户 id" json:"merchant_id"`           // 商户 id
	ProductID  int64                 `gorm:"column:product_id;not null;comment:产品 id" json:"product_id"`             // 产品 id
	AccountIds string                `gorm:"column:account_ids;not null;comment:子账户 id 拼接字符串" json:"account_ids"`    // 子账户 id 拼接字符串
	Status     int64                 `gorm:"column:status;not null;default:1;comment:状态:1 启用，2 关闭" json:"status"`    // 状态:1 启用，2 关闭
	CreateTime int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"` // 创建时间
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"` // 更新时间
	Delete_    soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`              // 删除状态:0 未删除，1 已删除
	Type       int64                 `gorm:"column:type;not null;comment:类型:1 代收，2 代付" json:"type"`                  // 类型:1 代收，2 代付
	Code       string                `gorm:"column:code;not null;comment:产品编码" json:"code"`                          // 产品编码
}

// TableName PayProductMerchantAccount's table name
func (*PayProductMerchantAccount) TableName() string {
	return TableNamePayProductMerchantAccount
}
