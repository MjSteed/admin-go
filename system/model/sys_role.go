package model

import "time"

// 角色表
type SysRole struct {
	Id         int64     `json:"id" form:"id" gorm:"primarykey"`       // 主键ID
	Name       string    `json:"name" form:"name" gorm:"comment:角色名称"` // 角色名称
	Code       string    `json:"code" form:"code" gorm:"comment:角色编码"` // 角色编码
	Sort       int       `json:"sort" form:"sort" gorm:"comment:显示顺序"`
	Status     int       `json:"status" form:"status" gorm:"comment:角色状态(1-正常；0-停用)"`             // 角色状态(1-正常；0-停用)
	Deleted    int       `json:"deleted" form:"deleted" gorm:"comment:逻辑删除标识(0-未删除；1-已删除)"`       // 逻辑删除标识(0-未删除；1-已删除)
	DataScope  int       `json:"dataScope" form:"dataScope" gorm:"comment:数据权限"`                  // 数据权限
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"autoCreateTime,comment:创建时间"` // 创建时间
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"autoUpdateTime,comment:更新时间"` // 更新时间
}

func (SysRole) TableName() string {
	return "sys_role"
}

// 用户和角色关联表
type SysUserRole struct {
	UserId int64 //用户ID
	RoleId int64 //角色ID
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}

// 角色和菜单关联表
type SysRoleMenu struct {
	RoleId int64 //角色ID
	MenuId int64 //用户ID
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
