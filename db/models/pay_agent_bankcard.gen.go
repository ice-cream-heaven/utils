// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"gorm.io/datatypes"
	"gorm.io/plugin/soft_delete"
)

const TableNamePayAgentBankcard = "pay_agent_bankcard"

// PayAgentBankcard mapped from table <pay_agent_bankcard>
type PayAgentBankcard struct {
	ID          int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`           // ID
	AgentID     int64                 `gorm:"column:agent_id;not null;comment:代理 id" json:"agent_id"`                 // 代理 id
	AccountName string                `gorm:"column:account_name;not null;comment:持卡人姓名" json:"account_name"`         // 持卡人姓名
	AccountNo   string                `gorm:"column:account_no;not null;comment:卡号" json:"account_no"`                // 卡号
	BankName    string                `gorm:"column:bank_name;not null;comment:银行名称" json:"bank_name"`                // 银行名称
	BankBranch  string                `gorm:"column:bank_branch;not null;comment:支行" json:"bank_branch"`              // 支行
	Province    string                `gorm:"column:province;comment:省" json:"province"`                              // 省
	City        string                `gorm:"column:city;not null;comment:市" json:"city"`                             // 市
	Extra       datatypes.JSON        `gorm:"column:extra;comment:自定义扩展信息" json:"extra"`                              // 自定义扩展信息
	Status      int64                 `gorm:"column:status;not null;default:1;comment:状态:1 启用，2 关闭" json:"status"`    // 状态:1 启用，2 关闭
	Remark      string                `gorm:"column:remark;not null;comment:备注" json:"remark"`                        // 备注
	CreateTime  int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"` // 创建时间
	UpdateTime  int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"` // 更新时间
	Delete_     soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`              // 删除状态:0 未删除，1 已删除
}

// TableName PayAgentBankcard's table name
func (*PayAgentBankcard) TableName() string {
	return TableNamePayAgentBankcard
}
