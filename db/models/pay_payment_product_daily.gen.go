// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayPaymentProductDaily = "pay_payment_product_daily"

// PayPaymentProductDaily mapped from table <pay_payment_product_daily>
type PayPaymentProductDaily struct {
	ID               int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                      // ID
	ProductID        int64                 `gorm:"column:product_id;not null;comment:产品 id" json:"product_id"`                        // 产品 id
	ProductAccountID int64                 `gorm:"column:product_account_id;not null;comment:产品账号 id" json:"product_account_id"`      // 产品账号 id
	Currency         string                `gorm:"column:currency;not null;comment:货币代码：三字母大写" json:"currency"`                       // 货币代码：三字母大写
	TotalCount       int64                 `gorm:"column:total_count;not null;comment:总笔数" json:"total_count"`                        // 总笔数
	TotalAmount      float64               `gorm:"column:total_amount;not null;default:0.00;comment:总金额" json:"total_amount"`         // 总金额
	SuccessCount     int64                 `gorm:"column:success_count;not null;comment:成功笔数" json:"success_count"`                   // 成功笔数
	SuccessAmount    float64               `gorm:"column:success_amount;not null;default:0.00;comment:成功金额" json:"success_amount"`    // 成功金额
	AgentAmount      float64               `gorm:"column:agent_amount;not null;default:0.00;comment:代理成功金额" json:"agent_amount"`      // 代理成功金额
	TradeCharge      float64               `gorm:"column:trade_charge;not null;default:0.0000;comment:交易手续费" json:"trade_charge"`     // 交易手续费
	AgentCharge      float64               `gorm:"column:agent_charge;not null;default:0.0000;comment:代理手续费" json:"agent_charge"`     // 代理手续费
	FailureCount     int64                 `gorm:"column:failure_count;not null;comment:失败笔数" json:"failure_count"`                   // 失败笔数
	FailureAmount    float64               `gorm:"column:failure_amount;not null;default:0.00;comment:失败金额" json:"failure_amount"`    // 失败金额
	RefundCount      int64                 `gorm:"column:refund_count;not null;comment:退款笔数" json:"refund_count"`                     // 退款笔数
	RefundAmount     float64               `gorm:"column:refund_amount;not null;default:0.00;comment:退款金额" json:"refund_amount"`      // 退款金额
	Date             string                `gorm:"column:date;not null;comment:日报本地化日期" json:"date"`                                  // 日报本地化日期
	StartTime        int64                 `gorm:"column:start_time;not null;comment:开始时间" json:"start_time"`                         // 开始时间
	EndTime          int64                 `gorm:"column:end_time;not null;comment:结束时间" json:"end_time"`                             // 结束时间
	CreateTime       int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`            // 创建时间
	UpdateTime       int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`            // 更新时间
	Delete_          soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`                         // 删除状态:0 未删除，1 已删除
	ChannelCharge    float64               `gorm:"column:channel_charge;not null;default:0.0000;comment:通道手续费" json:"channel_charge"` // 通道手续费
}

// TableName PayPaymentProductDaily's table name
func (*PayPaymentProductDaily) TableName() string {
	return TableNamePayPaymentProductDaily
}