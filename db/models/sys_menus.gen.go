// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import "gorm.io/plugin/soft_delete"

const TableNameSysMenu = "sys_menus"

// SysMenu mapped from table <sys_menus>
type SysMenu struct {
	ID                 int64                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreateTime         int64                 `gorm:"column:create_time;type:bigint;autoCreateTime:milli" json:"create_time"`                                                     // Create Time | 创建日期
	UpdateTime         int64                 `gorm:"column:update_time;type:bigint;autoUpdateTime:milli" json:"update_time"`                                                     // Update Time | 修改日期
	Sort               int64                 `gorm:"column:sort;not null;default:1;comment:Sort Number | 排序编号" json:"sort"`                                                      // Sort Number | 排序编号
	MenuLevel          int64                 `gorm:"column:menu_level;not null;comment:Menu level | 菜单层级" json:"menu_level"`                                                     // Menu level | 菜单层级
	MenuType           int64                 `gorm:"column:menu_type;not null;comment:Menu type | 菜单类型（菜单或目录）1 目录 2 菜单" json:"menu_type"`                                        // Menu type | 菜单类型（菜单或目录）1 目录 2 菜单
	Path               string                `gorm:"column:path;comment:Index path | 菜单路由路径" json:"path"`                                                                        // Index path | 菜单路由路径
	Name               string                `gorm:"column:name;not null;comment:Index name | 菜单名称" json:"name"`                                                                 // Index name | 菜单名称
	Redirect           string                `gorm:"column:redirect;comment:Redirect path | 跳转路径（外链）" json:"redirect"`                                                           // Redirect path | 跳转路径（外链）
	Component          string                `gorm:"column:component;comment:The path of vue file | 组件路径" json:"component"`                                                      // The path of vue file | 组件路径
	Disabled           int64                 `gorm:"column:disabled;comment:Disable status | 是否停用 1 停用 2 没停用" json:"disabled"`                                                   // Disable status | 是否停用 1 停用 2 没停用
	Title              string                `gorm:"column:title;not null;comment:Menu name | 菜单显示标题" json:"title"`                                                              // Menu name | 菜单显示标题
	Icon               string                `gorm:"column:icon;not null;comment:Menu icon | 菜单图标" json:"icon"`                                                                  // Menu icon | 菜单图标
	HideMenu           int64                 `gorm:"column:hide_menu;comment:Hide menu | 是否隐藏菜单 1 隐藏 2 不隐藏" json:"hide_menu"`                                                    // Hide menu | 是否隐藏菜单 1 隐藏 2 不隐藏
	HideBreadcrumb     int64                 `gorm:"column:hide_breadcrumb;comment:Hide the breadcrumb | 隐藏面包屑 1 隐藏 2 不隐藏" json:"hide_breadcrumb"`                               // Hide the breadcrumb | 隐藏面包屑 1 隐藏 2 不隐藏
	IgnoreKeepAlive    int64                 `gorm:"column:ignore_keep_alive;comment:Do not keep alive the tab | 取消页面缓存 1 取消缓存 2 不取消缓存" json:"ignore_keep_alive"`                // Do not keep alive the tab | 取消页面缓存 1 取消缓存 2 不取消缓存
	HideTab            int64                 `gorm:"column:hide_tab;comment:Hide the tab header | 隐藏页头 1 隐藏 2 不隐藏" json:"hide_tab"`                                              // Hide the tab header | 隐藏页头 1 隐藏 2 不隐藏
	FrameSrc           string                `gorm:"column:frame_src;comment:Show iframe | 内嵌 iframe" json:"frame_src"`                                                          // Show iframe | 内嵌 iframe
	CarryParam         int64                 `gorm:"column:carry_param;comment:The route carries parameters or not | 携带参数 1 带参数 2 不带参数" json:"carry_param"`                      // The route carries parameters or not | 携带参数 1 带参数 2 不带参数
	HideChildrenInMenu int64                 `gorm:"column:hide_children_in_menu;comment:Hide children menu or not | 隐藏所有子菜单 1 隐藏 2 不隐藏" json:"hide_children_in_menu"`           // Hide children menu or not | 隐藏所有子菜单 1 隐藏 2 不隐藏
	Affix              int64                 `gorm:"column:affix;comment:Affix tab | Tab 固定 1 固定 2 不固定" json:"affix"`                                                            // Affix tab | Tab 固定 1 固定 2 不固定
	DynamicLevel       int64                 `gorm:"column:dynamic_level;default:20;comment:The maximum number of pages the router can open | 能打开的子 TAB 数" json:"dynamic_level"` // The maximum number of pages the router can open | 能打开的子 TAB 数
	RealPath           string                `gorm:"column:real_path;comment:The real path of the route without dynamic part | 菜单路由不包含参数部分" json:"real_path"`                    // The real path of the route without dynamic part | 菜单路由不包含参数部分
	ParentID           int64                 `gorm:"column:parent_id;default:100000;comment:Parent menu ID | 父菜单 ID" json:"parent_id"`                                           // Parent menu ID | 父菜单 ID
	Delete_            soft_delete.DeletedAt `gorm:"column:_delete;not null;comment:删除状态:0 未删除，1 已删除" json:"-"`                                                                  // 删除状态:0 未删除，1 已删除
	Type               int64                 `gorm:"column:type;not null;default:1;comment:菜单类型 1: manager, 2: merchant, 3: agent" json:"type"`                                  // 菜单类型 1: manager, 2: merchant, 3: agent
}

// TableName SysMenu's table name
func (*SysMenu) TableName() string {
	return TableNameSysMenu
}