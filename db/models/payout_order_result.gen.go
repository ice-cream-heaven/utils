// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"gorm.io/datatypes"
	"gorm.io/plugin/soft_delete"
)

const TableNamePayoutOrderResult = "payout_order_result"

// PayoutOrderResult mapped from table <payout_order_result>
type PayoutOrderResult struct {
	ID            int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`           // ID
	OrderID       int64                 `gorm:"column:order_id;not null;comment:订单 id" json:"order_id"`                 // 订单 id
	OrderNo       string                `gorm:"column:order_no;not null;comment:交易订单号" json:"order_no"`                 // 交易订单号
	ChannelResult datatypes.JSON        `gorm:"column:channel_result;comment:上游数据返回" json:"channel_result"`             // 上游数据返回
	CreateTime    int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"` // 创建时间
	UpdateTime    int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"` // 更新时间
	Delete_       soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`              // 删除状态:0 未删除，1 已删除
}

// TableName PayoutOrderResult's table name
func (*PayoutOrderResult) TableName() string {
	return TableNamePayoutOrderResult
}
