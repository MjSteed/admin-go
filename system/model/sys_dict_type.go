package model

//字典类型表
type SysDictType struct {
	Id     int64  `gorm:"primarykey"`                                        // 主键ID
	Name   string `json:"name" form:"name" gorm:"comment:类型名称"`              // 类型名称
	Code   string `json:"code" form:"code" gorm:"comment:类型编码"`              // 类型编码
	Status int8   `json:"status" form:"status" gorm:"comment:状态(1:正常;0:禁用)"` // 状态(1:正常;0:禁用)
	Remark string `json:"remark" form:"remark" gorm:"comment:备注"`            //备注
}
