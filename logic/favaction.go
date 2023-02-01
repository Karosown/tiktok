package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"tiktok/dao"
	"tiktok/models"
	"tiktok/web"
)

type fres struct {
	Uid       int64  `gorm:"column:uid;primary_key;AUTO_INCREMENT" json:"Uid"` // 用户id
	Token     string `gorm:"column:token" json:"Token"`
	LoveVideo string `gorm:"column:love_video" json:"love_video"`
}
type vres struct {
	Vid           int64 `gorm:"column:id" json:"Vid"`                       // 视频唯一标识
	FavoriteCount int64 `gorm:"column:favorite_count" json:"FavoriteCount"` // 视频的点赞总数
	IsFavorite    bool  `gorm:"column:is_favorite" json:"IsFavorite"`       // true-已点赞，false-未点赞
}

func FavouriteAction(ctx *gin.Context) {
	token := ctx.Query("token")
	vid := ctx.Query("video_id")
	anctype := ctx.Query("action_type")
	db := dao.GetDB(ctx)
	tx := db.Begin()
	user := fres{}
	err := tx.Table("tiktok_user").Select("uid, token,love_video").Where("token = ?", token).Find(&user).Error
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

	vd := vres{}
	intvid, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		logrus.Error("vid转化失败")
	}
	err = tx.Table("tiktok_video").Where("vid=?", intvid).Find(&vd).Error
	if err != nil {
		logrus.Error("查询vid错误")
		ctx.JSON(http.StatusOK, models.Status{
			StatusCode: -1,
			StatusMsg:  "fail",
		})
		tx.Rollback()
		return
	}
	vd.IsFavorite = strings.Contains(user.LoveVideo, vid)
	if vd.IsFavorite != true && anctype == "1" {
		err = db.Table("tiktok_video").Model(&vres{Vid: vd.Vid}).Updates(map[string]interface{}{
			"favourite_count": vd.FavoriteCount + 1,
			"is_follow":       true,
			"love_video":      user.LoveVideo + vid + " ",
		}).Error
		if err != nil {

			logrus.Error("点赞发生错误")
			ctx.JSON(http.StatusOK, models.Status{
				StatusCode: -1,
				StatusMsg:  "fail",
			})
			tx.Rollback()
			return
		}
	}
	if vd.IsFavorite == true && anctype == "2" {
		result := strings.Replace(user.LoveVideo, vid, "", -1)
		err = db.Table("tiktok_video").Model(&vres{Vid: vd.Vid}).Updates(map[string]interface{}{
			"favourite_count": vd.FavoriteCount - 1,
			"is_follow":       false,
			"love_video":      result,
		}).Error
		if err != nil {
			logrus.Error("取消点赞发生错误")
			ctx.JSON(http.StatusOK, models.Status{
				StatusCode: -1,
				StatusMsg:  "fail",
			})
			tx.Rollback()
			return
		}
	}
	ctx.JSON(http.StatusOK, models.Status{
		StatusCode: 0,
		StatusMsg:  "favourite action success",
	})
}