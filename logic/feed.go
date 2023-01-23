package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Feed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "zhangsan",
	})

}
