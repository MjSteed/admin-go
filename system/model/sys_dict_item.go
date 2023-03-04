package model

//字典数据表
type SysDictItem struct {
	Id        int64  `gorm:"primarykey"`                                              // 主键ID
	TypeCode  string `json:"typeCode" form:"typeCode" gorm:"comment:字典类型编码"`          // 字典类型编码
	Name      string `json:"name" form:"name" gorm:"comment:字典项名称"`                   // 字典项名称
	Value     string `json:"value" form:"value" gorm:"comment:字典项值"`                  // 字典项值
	Sort      string `json:"sort" form:"sort" gorm:"comment:排序"`                      // 排序
	Status    int8   `json:"status" form:"status" gorm:"comment:状态(1:正常;0:禁用)"`       // 状态(1:正常;0:禁用)
	Defaulted int8   `json:"defaulted" form:"defaulted" gorm:"comment:是否默认(1:是;0:否)"` //是否默认(1:是;0:否)
	Remark    string `json:"remark" form:"remark" gorm:"comment:备注"`                  //备注
}

func (SysDictItem) TableName() string {
	return "sys_dict_item"
}
