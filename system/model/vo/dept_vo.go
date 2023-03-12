package vo

import (
	"time"

	"github.com/MjSteed/vue3-element-admin-go/system/model"
)

type Dept struct {
	Id         int64     `json:"id"`
	ParentId   int64     `json:"parentId"`
	Name       string    `json:"name"`
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	Children   []Dept    `json:"children"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func (dept *Dept) Format(d model.SysDept) (v Dept) {
	v.Id = d.Id
	v.ParentId = d.ParentId
	v.Name = d.Name
	v.Sort = d.Sort
	v.Status = d.Status
	v.CreateTime = d.CreateTime
	v.UpdateTime = d.UpdateTime
	return
}
