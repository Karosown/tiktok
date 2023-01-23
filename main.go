package main

import (
	"tiktok/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.RouterInit(r)
	r.Run()
}
