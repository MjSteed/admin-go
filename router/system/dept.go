package system

import (
	"github.com/MjSteed/vue3-element-admin-go/system/api"
	"github.com/gin-gonic/gin"
)

// DeptRouter 部门路由
type DeptRouter struct {
	api *api.DeptApi
}

// NewDeptRouter 创建部门路由
func NewDeptRouter(api *api.DeptApi) *DeptRouter {
	return &DeptRouter{api: api}
}

// InitRouter 初始化部门路由
func (rt *DeptRouter) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("/v1/dept")
	r.GET("", rt.api.ListPages)
	r.GET("/options", rt.api.ListOptions)
	r.GET("/:id/form", rt.api.GetForm)
	r.POST("", rt.api.Save)
	r.PUT("", rt.api.Update)
	r.DELETE("/:ids", rt.api.BatchDelete)
}
