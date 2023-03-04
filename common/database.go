package common

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func init() {
	mysqlConfig := mysql.Config{
		DSN: "root:root@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	}
	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	DB = db
}
