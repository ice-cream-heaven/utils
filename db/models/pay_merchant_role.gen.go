// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayMerchantRole = "pay_merchant_role"

// PayMerchantRole mapped from table <pay_merchant_role>
type PayMerchantRole struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`           // ID
	MerchantID int64                 `gorm:"column:merchant_id;comment:商户 id" json:"merchant_id"`                    // 商户 id
	Name       string                `gorm:"column:name;not null;comment:名称" json:"name"`                            // 名称
	Code       string                `gorm:"column:code;not null;comment:编码" json:"code"`                            // 编码
	Rules      string                `gorm:"column:rules;not null;comment:权限 ID 拼接字符串" json:"rules"`                 // 权限 ID 拼接字符串
	Sort       int64                 `gorm:"column:sort;not null;comment:自然数排序值" json:"sort"`                        // 自然数排序值
	Status     int64                 `gorm:"column:status;not null;default:1;comment:状态:1 启用，0 禁用" json:"status"`    // 状态:1 启用，0 禁用
	Remark     string                `gorm:"column:remark;not null;comment:备注" json:"remark"`                        // 备注
	CreateTime int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"` // 创建时间
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"` // 更新时间
	Default    int64                 `gorm:"column:_default;not null;comment:缺省状态:0 无，1 缺省｜不可删" json:"_default"`     // 缺省状态:0 无，1 缺省｜不可删
	Delete_    soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`              // 删除状态:0 未删除，1 已删除
}

// TableName PayMerchantRole's table name
func (*PayMerchantRole) TableName() string {
	return TableNamePayMerchantRole
}
