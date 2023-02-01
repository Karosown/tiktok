package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"tiktok/dao"
	"tiktok/models"
	"tiktok/web"
)

type flres struct {
	Status models.Status
	vdList []models.Video
}
type favuser struct {
	Uid       int64  `gorm:"column:uid;primary_key;AUTO_INCREMENT" json:"Uid"` // 用户id
	Token     string `gorm:"column:token" json:"Token"`
	LoveVideo string `gorm:"column:love_video" json:"love_video"`
}

func FavList(ctx *gin.Context) {
	uid := ctx.Query("user_id")
	token := ctx.Query("token")
	db := dao.GetDB(ctx)
	tx := db.Begin()
	user := favuser{}
	reuid, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		logrus.Error("information uid字符串转换失败")
	}
	err = tx.Table("tiktok_user").Select("uid, token,love_video").Where("uid = ?", reuid).Find(&user).Error
	if err != nil {
		logrus.Error("查询错误")
		ctx.JSON(http.StatusOK, models.Status{
			StatusCode: -1,
			StatusMsg:  "fail",
		})
		tx.Rollback()
		return
	}
	claim, err := web.ParseToken(token)
	if err != nil {
		retoken, err := web.CreateToken(claim.Uid, claim.Username)
		if err != nil {
			logrus.Error("重发token失败")
			ctx.JSON(http.StatusOK, models.Status{
				StatusCode: -1,
				StatusMsg:  "fail",
			})
			tx.Rollback()
			return
		}

		err = tx.Table("tiktok_user").Where("uer_id=?", claim.Uid).Update("token", retoken).Error
		if err != nil {
			logrus.Error("token写回db失败")
			ctx.JSON(http.StatusOK, models.Status{
				StatusCode: -1,
				StatusMsg:  "fail",
			})
			tx.Rollback()
			return
		}
	}
	lovelist := user.LoveVideo

	var videos []models.Video
	limit := 10
	page := 1
	offset := (page - 1) * limit
	db.Limit(limit).Offset(offset).Find(&videos)

	ctx.JSON(http.StatusOK, flres{
		Status: models.Status{
			StatusCode: 0,
			StatusMsg:  "favlist success",
		},
		vdList: videos,
	})
}
