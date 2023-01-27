package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/web"
)

func main() {
	r := gin.Default()
	web.RouterInit(r)
	//gin.SetMode(gin.ReleaseMode)
	r.Run() //127.0.0.1:8080
}
