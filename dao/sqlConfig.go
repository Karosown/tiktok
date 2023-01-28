package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Database *database

func DBConfig() *gorm.DB {
	c := Config{}
	c = *InitConfig()
	Database = &c.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", Database.UserName, Database.Password, Database.Host, Database.Port, Database.DbName, Database.ConnMaxLifetime)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tiktok_", // 表名前缀
			SingularTable: true,      // 单数表名
			NoLowerCase:   false,     // 关闭小写转换
		},
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	return db
}
