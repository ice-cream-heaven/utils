// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"gorm.io/datatypes"
	"gorm.io/plugin/soft_delete"
)

const TableNamePaymentOrderArchive230405013425 = "payment_order_archive_230405013425"

// PaymentOrderArchive230405013425 mapped from table <payment_order_archive_230405013425>
type PaymentOrderArchive230405013425 struct {
	ID                  int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                                                 // ID
	MerchantID          int64                 `gorm:"column:merchant_id;not null;comment:商户 id" json:"merchant_id"`                                                 // 商户 id
	SupplierID          int64                 `gorm:"column:supplier_id;comment:供应商 id" json:"supplier_id"`                                                         // 供应商 id
	AgentID             int64                 `gorm:"column:agent_id;comment:商户&代理 id" json:"agent_id"`                                                             // 商户&代理 id
	OrderNo             string                `gorm:"column:order_no;not null;comment:交易订单号" json:"order_no"`                                                       // 交易订单号
	MerchantOrderNo     string                `gorm:"column:merchant_order_no;not null;comment:商户订单号" json:"merchant_order_no"`                                     // 商户订单号
	ChannelOrderNo      string                `gorm:"column:channel_order_no;not null;comment:渠道订单号" json:"channel_order_no"`                                       // 渠道订单号
	ChannelTradeNo      string                `gorm:"column:channel_trade_no;not null;comment:渠道用户交易号" json:"channel_trade_no"`                                     // 渠道用户交易号
	ProductID           int64                 `gorm:"column:product_id;not null;comment:产品 id" json:"product_id"`                                                   // 产品 id
	ProductAccountID    int64                 `gorm:"column:product_account_id;not null;comment:产品资源 id" json:"product_account_id"`                                 // 产品资源 id
	Amount              float64               `gorm:"column:amount;not null;default:0.00;comment:交易金额" json:"amount"`                                               // 交易金额
	Country             string                `gorm:"column:country;comment:国家代码，中国:CN" json:"country"`                                                             // 国家代码，中国:CN
	Currency            string                `gorm:"column:currency;not null;default:CNY;comment:三位货币代码，人民币:CNY" json:"currency"`                                  // 三位货币代码，人民币:CNY
	ChannelRate         float64               `gorm:"column:channel_rate;not null;default:0.000000;comment:渠道费率" json:"channel_rate"`                               // 渠道费率
	ChannelCharge       float64               `gorm:"column:channel_charge;not null;default:0.00;comment:渠道费用" json:"channel_charge"`                               // 渠道费用
	TradeRate           float64               `gorm:"column:trade_rate;not null;default:0.000000;comment:商户费率" json:"trade_rate"`                                   // 商户费率
	TradeCharge         float64               `gorm:"column:trade_charge;not null;default:0.00;comment:商户费用" json:"trade_charge"`                                   // 商户费用
	AgentRate           float64               `gorm:"column:agent_rate;not null;default:0.000000;comment:代理费率" json:"agent_rate"`                                   // 代理费率
	AgentCharge         float64               `gorm:"column:agent_charge;not null;default:0.00;comment:代理费用" json:"agent_charge"`                                   // 代理费用
	Status              int64                 `gorm:"column:status;not null;comment:状态:0 生成订单，1 支付中，2 支付未通知，3 支付已通知，-1 交易失败，-2 交易过期，-3 交易撤销，-4 交易异常" json:"status"` // 状态:0 生成订单，1 支付中，2 支付未通知，3 支付已通知，-1 交易失败，-2 交易过期，-3 交易撤销，-4 交易异常
	StatusChange        string                `gorm:"column:status_change;comment:状态变化记录" json:"status_change"`                                                     // 状态变化记录
	Way                 int64                 `gorm:"column:way;not null;default:1;comment:处理方式：1 自动，2 手动" json:"way"`                                              // 处理方式：1 自动，2 手动
	ReturnURL           string                `gorm:"column:return_url;comment:同步回调 url" json:"return_url"`                                                         // 同步回调 url
	NotifyURL           string                `gorm:"column:notify_url;comment:异步回调 url" json:"notify_url"`                                                         // 异步回调 url
	ReceiptURL          string                `gorm:"column:receipt_url;comment:收据 url" json:"receipt_url"`                                                         // 收据 url
	Attach              string                `gorm:"column:attach;comment:附加参数" json:"attach"`                                                                     // 附加参数
	Result              datatypes.JSON        `gorm:"column:result;comment:数据返回" json:"result"`                                                                     // 数据返回
	ChannelErrorCode    string                `gorm:"column:channel_error_code;comment:上游支付错误码" json:"channel_error_code"`                                          // 上游支付错误码
	ChannelErrorMessage string                `gorm:"column:channel_error_message;comment:上游支付错误描述" json:"channel_error_message"`                                   // 上游支付错误描述
	Extra               datatypes.JSON        `gorm:"column:extra;comment:自定义扩展信息" json:"extra"`                                                                    // 自定义扩展信息
	ExpireTime          int64                 `gorm:"column:expire_time;comment:过期时间" json:"expire_time"`                                                           // 过期时间
	PayTime             int64                 `gorm:"column:pay_time;comment:支付时间" json:"pay_time"`                                                                 // 支付时间
	RefundTime          int64                 `gorm:"column:refund_time;comment:退款时间" json:"refund_time"`                                                           // 退款时间
	Remark              string                `gorm:"column:remark;not null;comment:备注" json:"remark"`                                                              // 备注
	CreateTime          int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`                                       // 创建时间
	UpdateTime          int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`                                       // 更新时间
	Delete_             soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`                                                    // 删除状态:0 未删除，1 已删除
}

// TableName PaymentOrderArchive230405013425's table name
func (*PaymentOrderArchive230405013425) TableName() string {
	return TableNamePaymentOrderArchive230405013425
}
