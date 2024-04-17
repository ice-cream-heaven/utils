// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"gorm.io/datatypes"
	"gorm.io/plugin/soft_delete"
)

const TableNamePaySupplierRechargeOrder = "pay_supplier_recharge_order"

// PaySupplierRechargeOrder mapped from table <pay_supplier_recharge_order>
type PaySupplierRechargeOrder struct {
	ID               int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                         // ID
	OrderNo          string                `gorm:"column:order_no;not null;comment:交易订单号" json:"order_no"`                               // 交易订单号
	SupplierID       int64                 `gorm:"column:supplier_id;not null;comment:供应商 id" json:"supplier_id"`                        // 供应商 id
	SystemBankcardID int64                 `gorm:"column:system_bankcard_id;not null;comment:系统收款卡 id" json:"system_bankcard_id"`        // 系统收款卡 id
	Amount           float64               `gorm:"column:amount;not null;default:0.00;comment:充值金额" json:"amount"`                       // 充值金额
	Currency         string                `gorm:"column:currency;not null;default:CNY;comment:三位货币代码，人民币:CNY" json:"currency"`          // 三位货币代码，人民币:CNY
	ExchangeCurrency string                `gorm:"column:exchange_currency;not null;comment:换汇货币代码，人民币:CNY" json:"exchange_currency"`    // 换汇货币代码，人民币:CNY
	ExchangeAmount   float64               `gorm:"column:exchange_amount;not null;default:0.0000;comment:换汇货币金额" json:"exchange_amount"` // 换汇货币金额
	IncomeAmount     float64               `gorm:"column:income_amount;not null;default:0.00;comment:到账金额" json:"income_amount"`         // 到账金额
	TradeCharge      float64               `gorm:"column:trade_charge;not null;default:0.00;comment:手续费" json:"trade_charge"`            // 手续费
	AccountName      string                `gorm:"column:account_name;not null;comment:持卡人姓名" json:"account_name"`                       // 持卡人姓名
	AccountNo        string                `gorm:"column:account_no;not null;comment:卡号" json:"account_no"`                              // 卡号
	BankName         string                `gorm:"column:bank_name;not null;comment:银行名称" json:"bank_name"`                              // 银行名称
	BankBranch       string                `gorm:"column:bank_branch;not null;comment:支行" json:"bank_branch"`                            // 支行
	Province         string                `gorm:"column:province;comment:省" json:"province"`                                            // 省
	City             string                `gorm:"column:city;not null;comment:市" json:"city"`                                           // 市
	Extra            datatypes.JSON        `gorm:"column:extra;comment:自定义扩展信息" json:"extra"`                                            // 自定义扩展信息
	Status           int64                 `gorm:"column:status;not null;comment:状态:0 待支付，1 已支付，2 已完成，-1 过期，-2 被驳回" json:"status"`       // 状态:0 待支付，1 已支付，2 已完成，-1 过期，-2 被驳回
	Remark           string                `gorm:"column:remark;not null;comment:备注" json:"remark"`                                      // 备注
	AuditFeedback    string                `gorm:"column:audit_feedback;comment:审核反馈" json:"audit_feedback"`                             // 审核反馈
	Auditor          int64                 `gorm:"column:auditor;comment:审核员 ID" json:"auditor"`                                         // 审核员 ID
	CreateTime       int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`               // 创建时间
	PayTime          int64                 `gorm:"column:pay_time;comment:支付时间" json:"pay_time"`                                         // 支付时间
	UpdateTime       int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`               // 更新时间
	Delete_          soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`                            // 删除状态:0 未删除，1 已删除
}

// TableName PaySupplierRechargeOrder's table name
func (*PaySupplierRechargeOrder) TableName() string {
	return TableNamePaySupplierRechargeOrder
}