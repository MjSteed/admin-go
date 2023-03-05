package model

import "time"

// 部门表
type SysDept struct {
	Id         int64     `gorm:"primarykey"`                                                // 主键ID
	Name       string    `json:"name" form:"name" gorm:"comment:部门名称"`                      // 部门名称
	ParentId   int64     `json:"parentId" form:"parentId" gorm:"comment:父节点id"`             // 父节点id
	TreePath   string    `json:"treePath" form:"treePath" gorm:"comment:父节点id路径"`           // 父节点id路径
	Sort       int       `json:"sort" form:"sort" gorm:"comment:显示顺序"`                      // 显示顺序
	Status     int       `json:"status" form:"status" gorm:"comment:状态(1:正常;0:禁用)"`         // 状态(1:正常;0:禁用)
	Deleted    int       `json:"deleted" form:"deleted" gorm:"comment:逻辑删除标识(1:已删除;0:未删除)"` // 逻辑删除标识(1:已删除;0:未删除)
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"comment:创建时间"`          // 创建时间
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"comment:更新时间"`          // 更新时间
}

func (SysDept) TableName() string {
	return "sys_dept"
}
