package vo

import (
	"time"

	"github.com/MjSteed/vue3-element-admin-go/system/model"
)

type Menu struct {
	Id         int64     `json:"id"`
	ParentId   int64     `json:"parentId"`
	Name       string    `json:"name"`
	Icon       string    `json:"icon"`
	RouteName  string    `json:"routeName"`
	RoutePath  string    `json:"routePath"`
	Component  string    `json:"component"`
	Redirect   string    `json:"redirect"`
	Sort       int       `json:"sort"`
	Visible    int       `json:"visible"`
	Type       string    `json:"type"`
	Perm       string    `json:"perm"`
	Children   []Menu    `json:"children"`
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
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
	v.Type = d.Type.String()
	v.Perm = d.Perm
	v.CreateTime = d.CreateTime
	v.UpdateTime = d.UpdateTime
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

// 获取菜单详情
type SysMenu struct {
	Id          int64     `json:"id"`          // 主键ID
	Name        string    `json:"name" `       // 菜单名称
	ParentId    int64     `json:"parentId" `   // 父节点id
	Type        string    `json:"type"`        // 菜单类型(1-菜单；2-目录；3-外链；4-按钮权限)
	Path        string    `json:"path"`        // 路由路径(浏览器地址栏路径)
	Component   string    `json:"component"`   // 组件路径(vue页面完整路径，省略.vue后缀)
	Perm        string    `json:"perm"`        // 权限标识
	Visible     int       `json:"visible"`     // 显示状态(1:显示;0:隐藏)
	Sort        int       `json:"sort"`        // 显示顺序
	Icon        string    `json:"icon"`        // 菜单图标
	RedirectUrl string    `json:"redirectUrl"` // 外链路径
	CreateTime  time.Time `json:"createTime"`  // 创建时间
	UpdateTime  time.Time `json:"updateTime"`  // 更新时间
}
