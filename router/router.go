package router

import (
	"net/http"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/middleware"
	"github.com/MjSteed/vue3-element-admin-go/system/api"
	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	r := gin.New()
	r.Use(middleware.GinLogger(common.LOG), middleware.GinRecovery(common.LOG, true))
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
		router.GET("/roles/pages", api.RoleApi.List)
		router.GET("/roles/options", api.RoleApi.ListOptions)
		router.GET("/roles/:id", api.RoleApi.GetForm)
		router.POST("/roles", api.RoleApi.Save)
		router.PUT("/roles/:id", api.RoleApi.Update)
		router.DELETE("/roles/:ids", api.RoleApi.BatchDelete)
		router.PUT("/roles/:id/status", api.RoleApi.UpdateRoleStatus)
		router.GET("/roles/:id/menuIds", api.RoleApi.GetRoleMenuIds)
		router.PUT("/roles/:id/menus", api.RoleApi.UpdateRoleMenus)

		router.GET("/users/me", api.UserApi.GetUserLoginInfo)
		router.GET("/users/pages", api.UserApi.List)
		router.GET("/users/:id/form", api.UserApi.GetForm)
		router.POST("/users", api.UserApi.Save)
		router.PUT("/users", api.UserApi.Update)
		router.DELETE("/users/:ids", api.UserApi.BatchDelete)
	}
	return r
}
