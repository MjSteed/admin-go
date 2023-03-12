package router

import (
	"net/http"

	"github.com/MjSteed/vue3-element-admin-go/middleware"
	"github.com/MjSteed/vue3-element-admin-go/system/api"
	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	r := gin.Default()

	pg := r.Group("test")
	{
		// 健康监测
		pg.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	r.Group("").POST("/api/v1/auth/login", api.AuthApi.Login)
	router := r.Group("/api/v1").Use(middleware.JWTAuth())
	// sr := system.DictItemRouter{}
	{
		router.GET("/dict/items/pages", api.DictItemApi.ListPages)
		router.GET("/dict/items/:id/form", api.DictItemApi.GetForm)
		router.POST("/dict/items", api.DictItemApi.Save)
		router.PUT("/dict/items/:id", api.DictItemApi.Update)
		router.DELETE("/dict/items/:ids", api.DictItemApi.BatchDelete)

		router.GET("/dict/types/pages", api.DictTypeApi.ListPages)
		router.GET("/dict/types/:id/form", api.DictTypeApi.GetForm)
		router.POST("/dict/types", api.DictTypeApi.Save)
		router.PUT("/dict/types/:id", api.DictTypeApi.Update)
		router.DELETE("/dict/types/:ids", api.DictTypeApi.BatchDelete)
		//路由冲突
		// r.GET("/:typeCode/items", api.ListDictItemsByTypeCode)

		router.GET("/dept", api.DeptApi.ListPages)
		router.GET("/dept/options", api.DeptApi.ListOptions)
		router.GET("/dept/:id/form", api.DeptApi.GetForm)
		router.POST("/dept", api.DeptApi.Save)
		router.PUT("/dept/:id", api.DeptApi.Update)
		router.DELETE("/dept/:ids", api.DeptApi.BatchDelete)

		menuApi := api.MenuApi{}
		router.GET("/menus/resources", menuApi.ListResources)
		router.GET("/menus", menuApi.List)
		router.GET("/menus/options", menuApi.ListOptions)
		router.GET("/menus/routes", menuApi.ListRoutes)
		router.POST("/menus", menuApi.Save)
		router.PUT("/menus", menuApi.Update)
		router.DELETE("/menus/:ids", menuApi.BatchDelete)

		roleApi := api.RoleApi{}
		router.GET("/roles/pages", roleApi.List)
		router.GET("/roles/options", roleApi.ListOptions)
		router.GET("/roles/:id", roleApi.GetForm)
		router.POST("/roles", roleApi.Save)
		router.PUT("/roles", roleApi.Update)
		router.DELETE("/roles/:ids", roleApi.BatchDelete)
		router.PUT("/roles/:id/status", roleApi.UpdateRoleStatus)
		router.GET("/roles/:id/menuIds", roleApi.GetRoleMenuIds)
		router.PUT("/roles/:id/menus", roleApi.UpdateRoleMenus)

		router.GET("/users/me", api.UserApi.GetUserLoginInfo)
		router.GET("/users/pages", api.UserApi.List)
		router.GET("/users/:id/form", api.UserApi.GetForm)
		router.POST("/users", api.UserApi.Save)
		router.PUT("/users", api.UserApi.Update)
		router.DELETE("/users/:ids", api.UserApi.BatchDelete)
	}
	return r
}
