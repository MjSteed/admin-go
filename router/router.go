package router

import (
	"net/http"

	"github.com/MjSteed/vue3-element-admin-go/router/system"
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
	api := r.Group("/api")
	sr := system.DictItemRouter{}
	{
		sr.InitDictItemRouter(api)
		sr.InitDictTypeRouter(api)
		sr.InitDeptRouter(api)
		sr.InitMenuRouter(api)
	}
	return r
}
