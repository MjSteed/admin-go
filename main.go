package main

import (
	"github.com/MjSteed/vue3-element-admin-go/bootstrap"
	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/router"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	bootstrap.InitializeConfig()
	common.LOG = bootstrap.InitLogger()
	common.Cache = bootstrap.InitializeRedis()
	common.DB = bootstrap.InitDatabase()
	r := router.Routers()
	r.Run(":9999")
}
