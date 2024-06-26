// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePaySupplierWithdrawlConfig = "pay_supplier_withdrawl_config"

// PaySupplierWithdrawlConfig mapped from table <pay_supplier_withdrawl_config>
type PaySupplierWithdrawlConfig struct {
	ID            int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                         // ID
	SupplierID    int64                 `gorm:"column:supplier_id;not null;comment:供应商 id" json:"supplier_id"`                        // 供应商 id
	TradeRate     float64               `gorm:"column:trade_rate;not null;default:0.000000;comment:充值手续费率" json:"trade_rate"`         // 充值手续费率
	ServiceCharge float64               `gorm:"column:service_charge;not null;default:0.00;comment:充值服务费" json:"service_charge"`      // 充值服务费
	StartTime     string                `gorm:"column:start_time;not null;comment:开始时间" json:"start_time"`                            // 开始时间
	EndTime       string                `gorm:"column:end_time;not null;comment:结束时间" json:"end_time"`                                // 结束时间
	UnitMinAmount float64               `gorm:"column:unit_min_amount;not null;default:0.00;comment:单笔最小交易金额" json:"unit_min_amount"` // 单笔最小交易金额
	UnitMaxAmount float64               `gorm:"column:unit_max_amount;not null;default:0.00;comment:单笔最大交易金额" json:"unit_max_amount"` // 单笔最大交易金额
	DayMaxCount   int64                 `gorm:"column:day_max_count;not null;comment:单日最大交易笔数" json:"day_max_count"`                  // 单日最大交易笔数
	DayMaxAmount  float64               `gorm:"column:day_max_amount;not null;default:0.00;comment:单日最大交易金额" json:"day_max_amount"`   // 单日最大交易金额
	SettleMethod  int64                 `gorm:"column:settle_method;not null;default:1;comment:结算周期：1T0，2T1" json:"settle_method"`    // 结算周期：1T0，2T1
	Status        int64                 `gorm:"column:status;not null;default:1;comment:状态:1 启用，2 关闭" json:"status"`                  // 状态:1 启用，2 关闭
	IsSystem      int64                 `gorm:"column:is_system;not null;comment:是否系统设置:1 是，0 不是" json:"is_system"`                   // 是否系统设置:1 是，0 不是
	CreateTime    int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`               // 创建时间
	UpdateTime    int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`               // 更新时间
	Delete_       soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`                            // 删除状态:0 未删除，1 已删除
}

// TableName PaySupplierWithdrawlConfig's table name
func (*PaySupplierWithdrawlConfig) TableName() string {
	return TableNamePaySupplierWithdrawlConfig
}
