package model

import "time"

// 字典类型表
type SysDictType struct {
	Id         int64     `json:"id" form:"id" gorm:"primarykey"`                                  // 主键ID
	Name       string    `json:"name" form:"name" gorm:"comment:类型名称"`                            // 类型名称
	Code       string    `json:"code" form:"code" gorm:"comment:类型编码"`                            // 类型编码
	Status     int8      `json:"status" form:"status" gorm:"comment:状态(1:正常;0:禁用)"`               // 状态(1:正常;0:禁用)
	Remark     string    `json:"remark" form:"remark" gorm:"comment:备注"`                          //备注
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"autoCreateTime;comment:创建时间"` // 创建时间
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"autoUpdateTime;comment:更新时间"` // 更新时间
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}
