package model

import "time"

// 字典数据表
type SysDictItem struct {
	Id         int64     `json:"id" form:"id" gorm:"primarykey"`                                  // 主键ID
	TypeCode   string    `json:"typeCode" form:"typeCode" gorm:"comment:字典类型编码"`                  // 字典类型编码
	Name       string    `json:"name" form:"name" gorm:"comment:字典项名称"`                           // 字典项名称
	Value      string    `json:"value" form:"value" gorm:"comment:字典项值"`                          // 字典项值
	Sort       int       `json:"sort" form:"sort" gorm:"comment:排序"`                              // 排序
	Status     int8      `json:"status" form:"status" gorm:"comment:状态(1:正常;0:禁用)"`               // 状态(1:正常;0:禁用)
	Defaulted  int8      `json:"defaulted" form:"defaulted" gorm:"comment:是否默认(1:是;0:否)"`         //是否默认(1:是;0:否)
	Remark     string    `json:"remark" form:"remark" gorm:"comment:备注"`                          //备注
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"autoCreateTime;comment:创建时间"` // 创建时间
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"autoUpdateTime;comment:更新时间"` // 更新时间
}

func (SysDictItem) TableName() string {
	return "sys_dict_item"
}
