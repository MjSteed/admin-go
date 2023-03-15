package bootstrap

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database 在中间件中初始化mysql链接
func InitDatabase() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN: "root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	}
	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		panic(fmt.Errorf("connnect to database failed: %s ", err))
	}
	return db
}
