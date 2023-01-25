package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Feed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
		"user_id":     123,
		"token":       "",
	})
}
