// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameRoleMenu = "role_menus"

// RoleMenu mapped from table <role_menus>
type RoleMenu struct {
	ID     int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoleID int64 `gorm:"column:role_id;not null" json:"role_id"`
	MenuID int64 `gorm:"column:menu_id;not null" json:"menu_id"`
}

// TableName RoleMenu's table name
func (*RoleMenu) TableName() string {
	return TableNameRoleMenu
}
