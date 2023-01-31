package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"tiktok/dao"
	"tiktok/models"
	"tiktok/web"
)

func (p Information) TableName() string {
	return "TIKTOK_USER"
}

type informationresponse struct {
	Status models.Status
	User   models.User
}
type Information struct {
	Uid           int64  `gorm:"column:uid;primary_key;AUTO_INCREMENT" json:"Uid"` // 用户id
	Name          string `gorm:"column:name" json:"Name"`                          // 用户名称
	FollowCount   int64  `gorm:"column:follow_count" json:"FollowCount"`           // 关注总数
	FollowerCount int64  `gorm:"column:follower_count" json:"FollowerCount"`       // 粉丝总数
	Token         string `gorm:"column:token" json:"Token"`
	IsFollow      bool   `gorm:"column:is_follow" json:"IsFollow"` // true-已关注，false-未关注
}

func Informatin(ctx *gin.Context) {
	uid := ctx.Query("user_id")
	token := ctx.Query("token")
	db := dao.GetDB(ctx)
	tx := db.Begin()
	user := Information{}
	result := tx.Where("uid=?", uid).Find(&user)
	if result.Error != nil {
		logrus.Error("注册查询信息失败")
		ctx.JSON(http.StatusBadRequest, informationresponse{
			Status: models.Status{
				StatusCode: -1,
				StatusMsg:  "注册返回信息失败",
			},
		})
	}
	_, err := web.ParseToken(token)
	if err != nil {
		restoken, err := web.CreateToken(user.Uid, user.Name)
		if err != nil {
			logrus.Error("重发token失败")
			ctx.JSON(http.StatusBadRequest, informationresponse{
				Status: models.Status{
					StatusCode: -1,
					StatusMsg:  "fail",
				},
			})
			tx.Rollback()
			return
		}
		err = tx.Where("uid = ?", uid).Update("token", restoken).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, informationresponse{
				Status: models.Status{
					StatusCode: -1,
					StatusMsg:  "注册返回信息失败",
				},
			})
			logrus.Error("login,db.Updata:更新token失败...")
			tx.Rollback()
			return
		}
	}
	ctx.JSON(http.StatusOK, informationresponse{
		Status: models.Status{
			StatusCode: 0,
			StatusMsg:  "注册返回信息成功",
		},
		User: models.User{
			Uid:           user.Uid,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		},
	})
	tx.Commit()
}
