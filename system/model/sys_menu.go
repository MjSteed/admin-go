package model

import "time"

// 菜单表
type SysMenu struct {
	Id          int64     `json:"id" form:"id" gorm:"primarykey"`                                      // 主键ID
	Name        string    `json:"name" form:"name" gorm:"comment:菜单名称"`                                // 菜单名称
	ParentId    int64     `json:"parentId" form:"parentId" gorm:"comment:父节点id"`                       // 父节点id
	Type        MenuType  `json:"type" form:"type" gorm:"comment:菜单类型(1-菜单；2-目录；3-外链；4-按钮权限)"`         // 菜单类型(1-菜单；2-目录；3-外链；4-按钮权限)
	Path        string    `json:"path" form:"path" gorm:"comment:路由路径(浏览器地址栏路径)"`                      // 路由路径(浏览器地址栏路径)
	Component   string    `json:"component" form:"component" gorm:"comment:组件路径(vue页面完整路径，省略.vue后缀)"`  // 组件路径(vue页面完整路径，省略.vue后缀)
	Perm        string    `json:"perm" form:"perm" gorm:"comment:权限标识"`                                // 权限标识
	Visible     int       `json:"visible" form:"visible" gorm:"comment:显示状态(1:显示;0:隐藏)"`               // 显示状态(1:显示;0:隐藏)
	Sort        int       `json:"sort" form:"sort" gorm:"comment:显示顺序"`                                // 显示顺序
	Icon        string    `json:"icon" form:"icon" gorm:"comment:菜单图标"`                                // 菜单图标
	RedirectUrl string    `json:"redirectUrl" form:"redirectUrl" gorm:"comment:外链路径"`                  // 外链路径
	CreateTime  time.Time `json:"createTime" form:"createTime" gorm:"autoCreateTime;comment:创建时间"`     // 创建时间
	UpdateTime  time.Time `json:"updateTime" form:"updateTime" gorm:"autoUpdateTime;comment:更新时间"`     // 更新时间
	SysRoles    []SysRole `gorm:"many2many:sys_role_menu;joinForeignKey:MenuId;joinReferences:RoleId"` //多对多关联
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

// 路由
type Route struct {
	Id          int64    `json:"id"`
	ParentId    int64    `json:"parentId"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Path        string   `json:"path"`
	Component   string   `json:"component"`
	Perm        string   `json:"perm"`
	Visible     int      `json:"visible"`
	Sort        int      `json:"sort"`
	Icon        string   `json:"icon"`
	RedirectUrl string   `json:"redirectUrl"`
	Roles       []string `json:"roles"`
}

// 菜单类型(1-菜单；2-目录；3-外链；4-按钮权限)
type MenuType int

const (
	//菜单
	MENU MenuType = iota + 1
	//目录
	CATALOG
	//外链
	EXTLINK
	//按钮
	BUTTON
)

// 获取菜单类型英文
func (t MenuType) String() string {
	menus := [...]string{"未知", "MENU", "CATALOG", "EXTLINK", "BUTTON"}
	if len(menus) < int(t) {
		return ""
	}
	return menus[t]
}
