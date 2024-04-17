// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayoutOrderConfirm = "payout_order_confirm"

// PayoutOrderConfirm mapped from table <payout_order_confirm>
type PayoutOrderConfirm struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	CreateTime int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`
	Delete_    soft_delete.DeletedAt `gorm:"column:_delete" json:"-"`
	OrderNo    string                `gorm:"column:order_no" json:"order_no"`
	UserID     int64                 `gorm:"column:user_id" json:"user_id"`
	Username   string                `gorm:"column:username" json:"username"`
	Action     string                `gorm:"column:action" json:"action"`
}

// TableName PayoutOrderConfirm's table name
func (*PayoutOrderConfirm) TableName() string {
	return TableNamePayoutOrderConfirm
}
