package common

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//用作参考，后面请注释掉
var R = gin.Default()

// Db 全局数据库操作实例
// 在mysql.go中初始化，并且在需要的时候返回
var Db *gorm.DB
