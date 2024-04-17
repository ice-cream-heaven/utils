// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayMerchantAccountChange = "pay_merchant_account_change"

// PayMerchantAccountChange mapped from table <pay_merchant_account_change>
type PayMerchantAccountChange struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                                             // ID
	MerchantID int64                 `gorm:"column:merchant_id;not null;comment:商户 id" json:"merchant_id"`                                             // 商户 id
	Currency   string                `gorm:"column:currency;not null;comment:货币代码：三字母大写" json:"currency"`                                              // 货币代码：三字母大写
	Balance    float64               `gorm:"column:balance;not null;default:0.0000;comment:可用余额快照" json:"balance"`                                     // 可用余额快照
	Deposit    float64               `gorm:"column:deposit;not null;default:0.0000;comment:代付余额快照" json:"deposit"`                                     // 代付余额快照
	Frozen     float64               `gorm:"column:frozen;not null;default:0.0000;comment:冻结快照" json:"frozen"`                                         // 冻结快照
	Profit     float64               `gorm:"column:profit;not null;default:0.0000;comment:收益快照" json:"profit"`                                         // 收益快照
	Amount     float64               `gorm:"column:amount;not null;default:0.0000;comment:变动金额" json:"amount"`                                         // 变动金额
	Source     int64                 `gorm:"column:source;not null;default:1;comment:变动来源：1 交易，2 充值，3 提现，4 推荐返佣，5 盈利转余额，6 余额转代付，7 投诉冻结" json:"source"` // 变动来源：1 交易，2 充值，3 提现，4 推荐返佣，5 盈利转余额，6 余额转代付，7 投诉冻结
	Type       int64                 `gorm:"column:type;not null;default:1;comment:变动类型：1 账户余额，2 盈利，3 代付余额" json:"type"`                               // 变动类型：1 账户余额，2 盈利，3 代付余额
	Way        int64                 `gorm:"column:way;not null;default:1;comment:变动方向：1 增加，2 减少" json:"way"`                                          // 变动方向：1 增加，2 减少
	Status     int64                 `gorm:"column:status;not null;default:1;comment:状态：1 冻结，2 已入账，3 已撤销" json:"status"`                               // 状态：1 冻结，2 已入账，3 已撤销
	OrderNo    string                `gorm:"column:order_no;not null;comment:关联系统订单号" json:"order_no"`                                                 // 关联系统订单号
	Remark     string                `gorm:"column:remark;not null;comment:备注" json:"remark"`                                                          // 备注
	CreateTime int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`                                   // 创建时间
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`                                   // 更新时间
	Delete_    soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`                                                // 删除状态:0 未删除，1 已删除
}

// TableName PayMerchantAccountChange's table name
func (*PayMerchantAccountChange) TableName() string {
	return TableNamePayMerchantAccountChange
}
