package router

import (
	"net/http"

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
	return r
}
