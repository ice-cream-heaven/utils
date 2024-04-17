// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNameMqQueueConfig = "mq_queue_config"

// MqQueueConfig mapped from table <mq_queue_config>
type MqQueueConfig struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreateTime int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`                                                                                     // Create Time | 创建日期
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`                                                                                     // Update Time | 修改日期
	Delete_    soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`                                                                                                  // 删除状态:0 未删除，1 已删除
	Status     int64                 `gorm:"column:status;default:1;comment:Status 1 : normal 2 : ban | 状态 1 正常 2 禁用" json:"status"`                                                                     // Status 1 : normal 2 : ban | 状态 1 正常 2 禁用
	Code       string                `gorm:"column:code;not null;comment:Queue Code | 队列编号，eg: 0001, 0002,0003" json:"code"`                                                                             // Queue Code | 队列编号，eg: 0001, 0002,0003
	QueueType  string                `gorm:"column:queue_type;not null;default:default;comment:queue type | 队列类型，default: 默认（老队列），shared: 共享（Hash 分流），exclusive（独享）" json:"queue_type"`                  // queue type | 队列类型，default: 默认（老队列），shared: 共享（Hash 分流），exclusive（独享）
	BizType    string                `gorm:"column:biz_type;not null;comment:Queue Business Type | 队列业务类型，payment: 代收，payout: 代付，order: 代收 + 代付" json:"biz_type"`                                        // Queue Business Type | 队列业务类型，payment: 代收，payout: 代付，order: 代收 + 代付
	TaskType   string                `gorm:"column:task_type;not null;comment:Queue Task Type | 任务类型，payout_async: 代付异步处理，settlement: 结算，channel_webhook: 上游回调，merchant_webhook: 商户回调" json:"task_type"` // Queue Task Type | 任务类型，payout_async: 代付异步处理，settlement: 结算，channel_webhook: 上游回调，merchant_webhook: 商户回调
	Name       string                `gorm:"column:name;not null;comment:Queue Name | 队列名称，{biz_type}.{task_type}.{queue_type}.{code}, eg: order.settlement.exclusive.0001" json:"name"`                 // Queue Name | 队列名称，{biz_type}.{task_type}.{queue_type}.{code}, eg: order.settlement.exclusive.0001
}

// TableName MqQueueConfig's table name
func (*MqQueueConfig) TableName() string {
	return TableNameMqQueueConfig
}
