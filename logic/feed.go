package logic

import (
	"fmt"
	"net/http"
	"tiktok/dao"
	"tiktok/models"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// test
func Feed(ctx *gin.Context) {
	logrus.Debug("check something")
	logrus.Error("hello,world")
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
		"user_id":     123,
		"token":       "",
	})

	//================================================
	db := dao.GetDB(ctx)
	var videos []models.Video

	// 一页多少条
	limit := 10
	// 第几页
	page := 1
	offset := (page - 1) * limit
	db.Limit(limit).Offset(offset).Find(&videos)
	fmt.Println(videos)

	//================================================

	//测试logrus
	logrus.Debug("check something")
	logrus.Error("hello,world")
}
