// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayoutArchiveInfo = "payout_archive_info"

// PayoutArchiveInfo mapped from table <payout_archive_info>
type PayoutArchiveInfo struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`           // ID
	TableName_ string                `gorm:"column:table_name;not null;comment:表名" json:"table_name"`                // 表名
	TableSize  int64                 `gorm:"column:table_size;not null;comment:表数据量大小" json:"table_size"`            // 表数据量大小
	StartTime  int64                 `gorm:"column:start_time;not null;comment:起始订单时间" json:"start_time"`            // 起始订单时间
	EndTime    int64                 `gorm:"column:end_time;not null;comment:截止订单时间" json:"end_time"`                // 截止订单时间
	StartID    int64                 `gorm:"column:start_id;comment:起始订单 id" json:"start_id"`                        // 起始订单 id
	EndID      int64                 `gorm:"column:end_id;comment:截止订单 id" json:"end_id"`                            // 截止订单 id
	Status     int64                 `gorm:"column:status;not null;default:2;comment:状态:1 活跃，2 已静态化" json:"status"`  // 状态:1 活跃，2 已静态化
	Remark     string                `gorm:"column:remark;not null;comment:备注" json:"remark"`                        // 备注
	CreateTime int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"` // 创建时间
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"` // 更新时间
	Delete_    soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`              // 删除状态:0 未删除，1 已删除
}

// TableName PayoutArchiveInfo's table name
func (*PayoutArchiveInfo) TableName() string {
	return TableNamePayoutArchiveInfo
}