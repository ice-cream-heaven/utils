// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNamePayAdminLoginLog = "pay_admin_login_log"

// PayAdminLoginLog mapped from table <pay_admin_login_log>
type PayAdminLoginLog struct {
	ID            int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`           // ID
	AdminID       int64                 `gorm:"column:admin_id;not null;comment:管理员 id" json:"admin_id"`                // 管理员 id
	Device        string                `gorm:"column:device;comment:登录设备" json:"device"`                               // 登录设备
	DeviceVersion string                `gorm:"column:device_version;comment:设备版本号" json:"device_version"`              // 设备版本号
	Os            string                `gorm:"column:os;comment:操作系统" json:"os"`                                       // 操作系统
	OsVersion     string                `gorm:"column:os_version;comment:系统版本号" json:"os_version"`                      // 系统版本号
	IP            string                `gorm:"column:ip;comment:登录 ip" json:"ip"`                                      // 登录 ip
	CreateTime    int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"` // 登录时间
	Delete_       soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`              // 删除状态:0 未删除，1 已删除
}

// TableName PayAdminLoginLog's table name
func (*PayAdminLoginLog) TableName() string {
	return TableNamePayAdminLoginLog
}
