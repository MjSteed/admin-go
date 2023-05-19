package system

import (
	"github.com/MjSteed/vue3-element-admin-go/system/api"
	"github.com/gin-gonic/gin"
)

// MenuRouter 菜单路由
type MenuRouter struct {
	api *api.MenuApi
}

// NewMenuRouter 创建菜单路由
func NewMenuRouter(api *api.MenuApi) *MenuRouter {
	return &MenuRouter{api: api}
}

// InitMenuRouter 初始化菜单路由
func (rt *MenuRouter) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("/v1/menus")
	r.GET("/resources", rt.api.ListResources)
	r.GET("", rt.api.List)
	r.GET("/options", rt.api.ListOptions)
	r.GET("/routes", rt.api.ListRoutes)
	r.GET("/:id", rt.api.GetById)
	r.POST("", rt.api.Save)
	r.PUT(":id", rt.api.Update)
	r.DELETE("/:ids", rt.api.BatchDelete)
}
