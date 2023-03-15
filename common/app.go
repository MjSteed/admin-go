package common

import (
	"github.com/MjSteed/vue3-element-admin-go/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	LOG   *zap.Logger
	Cache *redis.Client
	//配置项
	Config *config.ApplicationConfig
)
