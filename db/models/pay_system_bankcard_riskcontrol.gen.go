// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePaySystemBankcardRiskcontrol = "pay_system_bankcard_riskcontrol"

// PaySystemBankcardRiskcontrol mapped from table <pay_system_bankcard_riskcontrol>
type PaySystemBankcardRiskcontrol struct {
	ID            int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                         // ID
	BankcardID    int64                 `gorm:"column:bankcard_id;not null;comment:银行卡 ID" json:"bankcard_id"`                        // 银行卡 ID
	StartTime     string                `gorm:"column:start_time;not null;comment:开始时间" json:"start_time"`                            // 开始时间
	EndTime       string                `gorm:"column:end_time;not null;comment:结束时间" json:"end_time"`                                // 结束时间
	UnitMinAmount float64               `gorm:"column:unit_min_amount;not null;default:0.00;comment:单笔最小交易金额" json:"unit_min_amount"` // 单笔最小交易金额
	UnitMaxAmount float64               `gorm:"column:unit_max_amount;not null;default:0.00;comment:单笔最大交易金额" json:"unit_max_amount"` // 单笔最大交易金额
	DayMaxCount   int64                 `gorm:"column:day_max_count;not null;comment:单日最大交易笔数" json:"day_max_count"`                  // 单日最大交易笔数
	DayMaxAmount  float64               `gorm:"column:day_max_amount;not null;default:0.00;comment:单日最大交易金额" json:"day_max_amount"`   // 单日最大交易金额
	Status        int64                 `gorm:"column:status;not null;default:1;comment:状态:1 启用，2 关闭" json:"status"`                  // 状态:1 启用，2 关闭
	IsSystem      int64                 `gorm:"column:is_system;not null;comment:是否系统设置:1 是，0 不是" json:"is_system"`                   // 是否系统设置:1 是，0 不是
	CreateTime    int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`               // 创建时间
	UpdateTime    int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`               // 更新时间
	Delete_       soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`                            // 删除状态:0 未删除，1 已删除
}

// TableName PaySystemBankcardRiskcontrol's table name
func (*PaySystemBankcardRiskcontrol) TableName() string {
	return TableNamePaySystemBankcardRiskcontrol
}
