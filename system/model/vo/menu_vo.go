package vo

import (
	"github.com/MjSteed/vue3-element-admin-go/system/model"
)

type Menu struct {
	Id        int64  `json:"id"`
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	RouteName string `json:"routeName"`
	RoutePath string `json:"routePath"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Sort      int    `json:"sort"`
	Visible   int    `json:"visible"`
	Type      int    `json:"type"`
	Perm      string `json:"perm"`
	Children  []Menu `json:"children"`
}

func (m *Menu) Format(d model.SysMenu) (v Menu) {
	v.Id = d.Id
	v.ParentId = d.ParentId
	v.Name = d.Name
	v.Sort = d.Sort
	v.Icon = d.Icon
	v.Component = d.Component
	v.Redirect = d.RedirectUrl
	v.Visible = d.Visible
	v.Type = d.Type
	v.Perm = d.Perm
	return
}

// 菜单路由视图对象
type Route struct {
	Path      string  `json:"path"`
	Component string  `json:"component"`
	Redirect  string  `json:"redirect"`
	Name      string  `json:"name"`
	Meta      Meta    `json:"meta"`
	Children  []Route `json:"children"`
}

type Meta struct {
	Title      string   `json:"title"`
	Icon       string   `json:"icon"`
	Hidden     bool     `json:"hidden"`
	AlwaysShow bool     `json:"alwaysShow"`
	Roles      []string `json:"roles"`
	KeepAlive  bool     `json:"keepAlive"`
}
