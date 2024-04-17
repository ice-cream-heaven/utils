// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayTelegram = "pay_telegram"

// PayTelegram mapped from table <pay_telegram>
type PayTelegram struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	CreateTime int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`
	Delete_    soft_delete.DeletedAt `gorm:"column:_delete" json:"-"`
	MerchantID int64                 `gorm:"column:merchant_id" json:"merchant_id"`
	AppID      string                `gorm:"column:app_id" json:"app_id"`
	ChatID     int64                 `gorm:"column:chat_id" json:"chat_id"`
	ChatTitle  string                `gorm:"column:chat_title" json:"chat_title"`
	UserID     int64                 `gorm:"column:user_id" json:"user_id"`
	Username   string                `gorm:"column:username" json:"username"`
	Role       int64                 `gorm:"column:role" json:"role"`
	Operate    string                `gorm:"column:operate" json:"operate"`
}

// TableName PayTelegram's table name
func (*PayTelegram) TableName() string {
	return TableNamePayTelegram
}
