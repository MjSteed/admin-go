package model

import "time"

// 菜单表
type SysMenu struct {
	Id          int64     `gorm:"primarykey"`                                                         // 主键ID
	Name        string    `json:"name" form:"name" gorm:"comment:菜单名称"`                               // 菜单名称
	ParentId    int64     `json:"parentId" form:"parentId" gorm:"comment:父节点id"`                      // 父节点id
	Type        int       `json:"type" form:"type" gorm:"comment:菜单类型(1-菜单；2-目录；3-外链；4-按钮权限)"`        // 菜单类型(1-菜单；2-目录；3-外链；4-按钮权限)
	Path        string    `json:"path" form:"path" gorm:"comment:路由路径(浏览器地址栏路径)"`                     // 路由路径(浏览器地址栏路径)
	Component   string    `json:"component" form:"component" gorm:"comment:组件路径(vue页面完整路径，省略.vue后缀)"` // 组件路径(vue页面完整路径，省略.vue后缀)
	Perm        string    `json:"perm" form:"perm" gorm:"comment:权限标识"`                               // 权限标识
	Visible     int       `json:"visible" form:"visible" gorm:"comment:显示状态(1:显示;0:隐藏)"`              // 显示状态(1:显示;0:隐藏)
	Sort        int       `json:"sort" form:"sort" gorm:"comment:显示顺序"`                               // 显示顺序
	Icon        string    `json:"icon" form:"icon" gorm:"comment:菜单图标"`                               // 菜单图标
	RedirectUrl string    `json:"redirectUrl" form:"redirectUrl" gorm:"comment:外链路径"`                 // 外链路径
	CreateTime  time.Time `json:"createTime" form:"createTime" gorm:"comment:创建时间"`                   // 创建时间
	UpdateTime  time.Time `json:"updateTime" form:"updateTime" gorm:"comment:更新时间"`                   // 更新时间
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

// 路由
type Route struct {
	Id          int64
	ParentId    int64
	Name        string
	Type        int
	Path        string
	Component   string
	Perm        string
	Visible     int
	Sort        int
	Icon        string
	RedirectUrl string
	Roles       []string
}
