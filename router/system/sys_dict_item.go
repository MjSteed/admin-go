package system

import (
	"github.com/MjSteed/vue3-element-admin-go/system/api"
	"github.com/gin-gonic/gin"
)

type DictItemRouter struct{}

func (rt *DictItemRouter) InitDictItemRouter(Router *gin.RouterGroup) {
	r := Router.Group("/v1/dict/items")
	r.GET("/pages", api.DictItemApi.ListPages)
	r.GET("/:id/form", api.DictItemApi.GetForm)
	r.POST("", api.DictItemApi.Save)
	r.PUT("", api.DictItemApi.Update)
	r.DELETE("/:ids", api.DictItemApi.BatchDelete)
}

func (rt *DictItemRouter) InitDictTypeRouter(Router *gin.RouterGroup) {
	r := Router.Group("/v1/dict/types")
	r.GET("/pages", api.DictTypeApi.ListPages)
	r.GET("/:id/form", api.DictTypeApi.GetForm)
	r.POST("", api.DictTypeApi.Save)
	r.PUT("", api.DictTypeApi.Update)
	r.DELETE("/:ids", api.DictTypeApi.BatchDelete)
	//路由冲突
	// r.GET("/:typeCode/items", api.ListDictItemsByTypeCode)
}

func (rt *DictItemRouter) InitRolesRouter(Router *gin.RouterGroup) {
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

func (rt *DictItemRouter) InitUserRouter(Router *gin.RouterGroup) {
	r := Router.Group("/v1/users")
	r.GET("/pages", api.UserApi.List)
	r.GET("/:id/form", api.UserApi.GetForm)
	r.POST("", api.UserApi.Save)
	r.PUT("", api.UserApi.Update)
	r.DELETE("/:ids", api.UserApi.BatchDelete)
}
