package model

import "time"

// 角色表
type SysUser struct {
	Id         int64     `json:"id" form:"id" gorm:"primarykey"`                                  // 主键ID
	Username   string    `json:"username" form:"username" gorm:"comment:用户名"`                     // 用户名
	Nickname   string    `json:"nickname" form:"nickname" gorm:"comment:昵称"`                      // 昵称
	Gender     int       `json:"gender" form:"gender" gorm:"comment:性别((1:男;2:女))"`               //性别((1:男;2:女))
	Password   string    `json:"password" form:"password" gorm:"comment:密码"`                      // 密码
	DeptId     int64     `json:"deptId" form:"deptId" gorm:"comment:部门ID"`                        // 部门ID
	Avatar     string    `json:"avatar" form:"avatar" gorm:"comment:用户头像"`                        // 用户头像
	Mobile     string    `json:"mobile" form:"mobile" gorm:"comment:联系方式"`                        // 联系方式
	Status     int       `json:"status" form:"status" gorm:"comment:用户状态((1:正常;0:禁用))"`           // 用户状态((1:正常;0:禁用))
	Email      string    `json:"email" form:"email" gorm:"comment:用户邮箱"`                          // 用户邮箱
	Deleted    int       `json:"deleted" form:"deleted" gorm:"comment:逻辑删除标识(0:未删除;1:已删除)"`       // 逻辑删除标识(0:未删除;1:已删除)
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"autoCreateTime,comment:创建时间"` // 创建时间
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"autoUpdateTime,comment:更新时间"` // 更新时间
}

func (*SysUser) TableName() string {
	return "sys_user"
}
