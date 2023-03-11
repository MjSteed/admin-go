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
		itemApi := api.DictItemApi{}
		router.GET("/dict/items/pages", itemApi.ListPages)
		router.GET("/dict/items/:id/form", itemApi.GetForm)
		router.POST("/dict/items", itemApi.Save)
		router.PUT("/dict/items", itemApi.Update)
		router.DELETE("/dict/items/:ids", itemApi.BatchDelete)

		typeApi := api.DictTypeApi{}
		router.GET("/dict/types/pages", typeApi.ListPages)
		router.GET("/dict/types/:id/form", typeApi.GetForm)
		router.POST("/dict/types", typeApi.Save)
		router.PUT("/dict/types", typeApi.Update)
		router.DELETE("/dict/types/:ids", typeApi.BatchDelete)
		//路由冲突
		// r.GET("/:typeCode/items", api.ListDictItemsByTypeCode)

		deptApi := api.DeptApi{}
		router.GET("/dept", deptApi.ListPages)
		router.GET("/dept/options", deptApi.ListOptions)
		router.GET("/dept/:id/form", deptApi.GetForm)
		router.POST("/dept", deptApi.Save)
		router.PUT("/dept", deptApi.Update)
		router.DELETE("/dept/:ids", deptApi.BatchDelete)

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
