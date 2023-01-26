package main

import (
	"tiktok/web"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	web.RouterInit(r)
	r.Run() //127.0.0.1:8080
}
