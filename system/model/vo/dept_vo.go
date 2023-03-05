package vo

import (
	"time"

	"github.com/MjSteed/vue3-element-admin-go/system/model"
)

type Dept struct {
	Id         int64
	ParentId   int64
	Name       string
	Sort       int
	Status     int
	Children   []Dept
	CreateTime time.Time
	UpdateTime time.Time
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
