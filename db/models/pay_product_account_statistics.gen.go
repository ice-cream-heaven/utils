// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayProductAccountStatistic = "pay_product_account_statistics"

// PayProductAccountStatistic mapped from table <pay_product_account_statistics>
type PayProductAccountStatistic struct {
	ID                  int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                                      // ID
	ProductAccountID    int64                 `gorm:"column:product_account_id;not null;comment:支付产品账户 id" json:"product_account_id"`                    // 支付产品账户 id
	OrderAllCount       int64                 `gorm:"column:order_all_count;not null;comment:全部订单量" json:"order_all_count"`                              // 全部订单量
	OrderCompleteCount  int64                 `gorm:"column:order_complete_count;not null;comment:已完成订单量" json:"order_complete_count"`                   // 已完成订单量
	OrderAllAmount      float64               `gorm:"column:order_all_amount;not null;default:0.00;comment:总订单交易金额" json:"order_all_amount"`             // 总订单交易金额
	OrderCompleteAmount float64               `gorm:"column:order_complete_amount;not null;default:0.00;comment:已完成订单交易金额" json:"order_complete_amount"` // 已完成订单交易金额
	OrderFailureCount   int64                 `gorm:"column:order_failure_count;not null;comment:订单连续失败笔数" json:"order_failure_count"`                   // 订单连续失败笔数
	AverageHandleTime   int64                 `gorm:"column:average_handle_time;not null;comment:完成订单平均处理时间" json:"average_handle_time"`                 // 完成订单平均处理时间
	LastPayTime         int64                 `gorm:"column:last_pay_time;not null;comment:最后支付时间" json:"last_pay_time"`                                 // 最后支付时间
	LastNode            string                `gorm:"column:last_node;not null;comment:最后计算节点" json:"last_node"`                                         // 最后计算节点
	CreateTime          int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`                            // 创建时间
	UpdateTime          int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`                            // 更新时间
	Delete_             soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`                                         // 删除状态:0 未删除，1 已删除
}

// TableName PayProductAccountStatistic's table name
func (*PayProductAccountStatistic) TableName() string {
	return TableNamePayProductAccountStatistic
}
