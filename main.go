package main

import (
	"tiktok/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitServer(r)
	//gin.SetMode(gin.ReleaseMode)
	r.Run() //127.0.0.1:8080
}
