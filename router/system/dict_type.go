package system

import (
	"github.com/MjSteed/vue3-element-admin-go/system/api"
	"github.com/gin-gonic/gin"
)

// DictTypeRouter 字典类型路由
type DictTypeRouter struct {
	api *api.DictTypeApi
}

// NewDictTypeRouter 创建字典类型路由
func NewDictTypeRouter(api *api.DictTypeApi) *DictTypeRouter {
	return &DictTypeRouter{api: api}
}

// InitDictTypeRouter 初始化字典类型路由
func (rt *DictTypeRouter) InitDictTypeRouter(Router *gin.RouterGroup) {
	r := Router.Group("/v1/dict/types")
	r.GET("/pages", rt.api.ListPages)
	r.GET("/:id/form", rt.api.GetForm)
	r.POST("", rt.api.Create)
	r.PUT("/:id", rt.api.Update)
	r.DELETE("/:ids", rt.api.BatchDelete)
	//路由冲突
	// r.GET("/:typeCode/items", api.ListDictItemsByTypeCode)
}
