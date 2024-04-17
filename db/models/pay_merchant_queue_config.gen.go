// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayMerchantQueueConfig = "pay_merchant_queue_config"

// PayMerchantQueueConfig mapped from table <pay_merchant_queue_config>
type PayMerchantQueueConfig struct {
	ID              int64                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreateTime      int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`           // Create Time | 创建日期
	UpdateTime      int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`           // Update Time | 修改日期
	Delete_         soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`                        // 删除状态:0 未删除，1 已删除
	MqQueueConfigID int64                 `gorm:"column:mq_queue_config_id;not null;comment:MQ Queue ID" json:"mq_queue_config_id"` // MQ Queue ID
	MerchantID      int64                 `gorm:"column:merchant_id;not null;comment:Merchant ID" json:"merchant_id"`               // Merchant ID
}

// TableName PayMerchantQueueConfig's table name
func (*PayMerchantQueueConfig) TableName() string {
	return TableNamePayMerchantQueueConfig
}