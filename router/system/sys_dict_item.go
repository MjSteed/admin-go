package system

import (
	"github.com/MjSteed/vue3-element-admin-go/system/api"
	"github.com/gin-gonic/gin"
)

type DictItemRouter struct{}

func (rt *DictItemRouter) InitDictItemRouter(Router *gin.RouterGroup) {
	api := api.DictItemApi{}
	r := Router.Group("/v1/dict/items")
	r.GET("/pages", api.ListPages)
	r.GET("/:id/form", api.GetForm)
	r.POST("", api.Save)
	r.PUT("", api.Update)
	r.DELETE("/:ids", api.BatchDelete)
}

func (rt *DictItemRouter) InitDictTypeRouter(Router *gin.RouterGroup) {
	api := api.DictTypeApi{}
	r := Router.Group("/v1/dict/types")
	r.GET("/pages", api.ListPages)
	r.GET("/:id/form", api.GetForm)
	r.POST("", api.Save)
	r.PUT("", api.Update)
	r.DELETE("/:ids", api.BatchDelete)
	//路由冲突
	// r.GET("/:typeCode/items", api.ListDictItemsByTypeCode)
}

func (rt *DictItemRouter) InitDeptRouter(Router *gin.RouterGroup) {
	api := api.DeptApi{}
	r := Router.Group("/v1/dept")
	r.GET("", api.ListPages)
	r.GET("/options", api.ListOptions)
	r.GET("/:id/form", api.GetForm)
	r.POST("", api.Save)
	r.PUT("", api.Update)
	r.DELETE("/:ids", api.BatchDelete)
}

func (rt *DictItemRouter) InitMenuRouter(Router *gin.RouterGroup) {
	api := api.MenuApi{}
	r := Router.Group("/v1/menus")
	r.GET("/resources", api.ListResources)
	r.GET("", api.List)
	r.GET("/options", api.ListOptions)
	r.GET("/routes", api.ListRoutes)
	r.POST("", api.Save)
	r.PUT("", api.Update)
	r.DELETE("/:ids", api.BatchDelete)
}
