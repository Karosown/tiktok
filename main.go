package main

import (
	"tiktok/config"
	"tiktok/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.MysqlInit()
	r := gin.Default()
	routers.RouterInit(r)
	r.Run() //127.0.0.1:8080
}
