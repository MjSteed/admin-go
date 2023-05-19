package system

import (
	"github.com/MjSteed/vue3-element-admin-go/system/api"
	"github.com/gin-gonic/gin"
)

// DictItemRouter 字典数据项路由
type DictItemRouter struct {
	api *api.DictItemApi
}

// NewDictItemRouter 创建字典数据项路由
func NewDictItemRouter(api *api.DictItemApi) *DictItemRouter {
	return &DictItemRouter{api: api}
}

// InitDictItemRouter 初始化字典数据项路由
func (rt *DictItemRouter) InitDictItemRouter(Router *gin.RouterGroup) {
	r := Router.Group("/v1/dict/items")
	r.GET("/pages", rt.api.ListPages)
	r.GET("/:id/form", rt.api.GetForm)
	r.POST("", rt.api.Save)
	r.PUT("/:id", rt.api.Update)
	r.DELETE("/:ids", rt.api.BatchDelete)
}
