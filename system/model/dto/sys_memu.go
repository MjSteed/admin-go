package dto

import (
	"github.com/MjSteed/vue3-element-admin-go/system/model"
)

type MenuForm struct {
	Id          int64  `form:"id"`          // 主键ID
	Name        string `form:"name"`        // 菜单名称
	ParentId    int64  `form:"parentId"`    // 父节点id
	Type        string `form:"type"`        // 菜单类型(MENU-菜单；CATALOG-目录；EXTLINK-外链；BUTTON-按钮权限)
	Path        string `form:"path"`        // 路由路径(浏览器地址栏路径)
	Component   string `form:"component"`   // 组件路径(vue页面完整路径，省略.vue后缀)
	Perm        string `form:"perm"`        // 权限标识
	Visible     int    `form:"visible"`     // 显示状态(1:显示;0:隐藏)
	Sort        int    `form:"sort"`        // 显示顺序
	Icon        string `form:"icon"`        // 菜单图标
	RedirectUrl string `form:"redirectUrl"` // 外链路径
}

func (m *MenuForm) ToMenu() model.SysMenu {
	d := model.SysMenu{
		Id:          m.Id,
		Name:        m.Name,
		ParentId:    m.ParentId,
		Path:        m.Path,
		Component:   m.Component,
		Perm:        m.Perm,
		Sort:        m.Sort,
		Visible:     m.Visible,
		Icon:        m.Icon,
		RedirectUrl: m.RedirectUrl,
	}
	switch m.Type {
	case "MENU":
		d.Type = 1
	case "CATALOG":
		d.Type = 2
	case "EXTLINK":
		d.Type = 3
	case "BUTTON":
		d.Type = 4
	}
	return d
}
