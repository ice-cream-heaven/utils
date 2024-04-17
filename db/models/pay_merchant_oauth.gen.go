// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayMerchantOauth = "pay_merchant_oauth"

// PayMerchantOauth mapped from table <pay_merchant_oauth>
type PayMerchantOauth struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增 id" json:"id"`        // 自增 id
	MerchantID int64                 `gorm:"column:merchant_id;not null;comment:商户 id" json:"merchant_id"`           // 商户 id
	Type       int64                 `gorm:"column:type;not null;default:1;comment:类型:1 代收，2 代付" json:"type"`        // 类型:1 代收，2 代付
	AccessKey  string                `gorm:"column:access_key;not null;comment:Access Key" json:"access_key"`        // Access Key
	SecretKey  string                `gorm:"column:secret_key;not null;comment:Secret Key" json:"secret_key"`        // Secret Key
	CreateTime int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"` // 创建时间
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"` // 更新时间
	Delete_    soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态：0 未删除，1 已删除" json:"-"`              // 删除状态：0 未删除，1 已删除
}

// TableName PayMerchantOauth's table name
func (*PayMerchantOauth) TableName() string {
	return TableNamePayMerchantOauth
}
