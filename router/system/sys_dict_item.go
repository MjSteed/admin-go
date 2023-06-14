package system

import (
	"github.com/MjSteed/vue3-element-admin-go/system/api"
	"github.com/gin-gonic/gin"
)

type dictItemRouter struct{}

func (rt *dictItemRouter) InitRolesRouter(Router *gin.RouterGroup) {
	r := Router.Group("/v1/roles")
	r.GET("/pages", api.RoleApi.List)
	r.GET("/options", api.RoleApi.ListOptions)
	r.GET("/:id", api.RoleApi.GetForm)
	r.POST("", api.RoleApi.Save)
	r.PUT("", api.RoleApi.Update)
	r.DELETE("/:ids", api.RoleApi.BatchDelete)
	r.PUT("/:id/status", api.RoleApi.UpdateRoleStatus)
	r.GET("/:id/menuIds", api.RoleApi.GetRoleMenuIds)
	r.PUT("/:id/menus", api.RoleApi.UpdateRoleMenus)
}

func (rt *dictItemRouter) InitUserRouter(Router *gin.RouterGroup) {
	r := Router.Group("/v1/users")
	r.GET("/pages", api.UserApi.List)
	r.GET("/:id/form", api.UserApi.GetForm)
	r.POST("", api.UserApi.Save)
	r.PUT("", api.UserApi.Update)
	r.DELETE("/:ids", api.UserApi.BatchDelete)
}
