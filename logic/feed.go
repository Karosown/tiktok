package logic

import (
	"net/http"
	"tiktok/models"

	"github.com/gin-gonic/gin"
)

// test
func Feed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
		"user_id":     123,
		"token":       "",
	})
}

// test
func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "string",
		"user": models.User{
			Uid:           0,
			Name:          "string",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      true,
		},
	})
}
