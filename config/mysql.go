package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"tiktok/common"
)

func MysqlInit() {
	// 日志打印
	newLogger := logger.Default

	log.Println("Mysql:初始化！")
	dsn := "root:你的密码@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	v, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // true不在表后面+ s，
		},
	})
	if err != nil {
		log.Panic(err)
	}
	common.Db = v
}
