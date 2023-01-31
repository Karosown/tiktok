package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"tiktok/dao"
	"tiktok/models"
	"tiktok/web"
)

func (p insert) TableName() string {
	return "TIKTOK_USER"
}

type registeresponse struct {
	StatusCode int64  `json:"StatusCode"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"StatusMsg"`  // 返回状态描述
	Token      string `json:"token"`      // 用户鉴权token
	UserID     int64  `json:"user_id"`    // 用户id
}
type insert struct {
	Uid           int64  `gorm:"column:uid;primary_key;AUTO_INCREMENT" json:"Uid"` // 用户id
	Name          string `gorm:"column:name" json:"Name"`                          // 用户名称
	FollowCount   int64  `gorm:"column:follow_count" json:"FollowCount"`           // 关注总数
	FollowerCount int64  `gorm:"column:follower_count" json:"FollowerCount"`       // 粉丝总数
	IsFollow      bool   `gorm:"column:is_follow" json:"IsFollow"`                 // true-已关注，false-未关注
	Password      string `gorm:"column:password" json:"PassWord"`
	Token         string `gorm:"column:token" json:"Token"`
}

func Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	db := dao.GetDB(ctx)

	tx := db.Begin()
	result := tx.Where("Name=?", username).Find(&models.User{})
	if result.RowsAffected == 0 {
		p := &insert{
			Name: username, FollowCount: 0, FollowerCount: 0, IsFollow: false, Password: password,
		}
		res := tx.Create(p)
		if res.Error != nil {
			ctx.JSON(http.StatusBadRequest, informationresponse{
				Status: models.Status{
					StatusCode: -1,
					StatusMsg:  "失败",
				},
			})
			logrus.Error("插入数据失败")
			tx.Rollback()
			return
		}
		uid := p.Uid
		token, err := web.CreateToken(uid, username)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, registeresponse{
				StatusCode: -1, StatusMsg: "fail", Token: "", UserID: 0,
			})

			logrus.Error("built token error")
			tx.Rollback()
			return
		}
		err = tx.Model(&insert{
			Uid: uid,
		}).Where("Name=?", username).Update("Token", token).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, registeresponse{
				StatusCode: -1, StatusMsg: "fail", Token: "", UserID: 0,
			})
			logrus.Error("更新token失败")
		}
		ctx.JSON(http.StatusOK, registeresponse{
			StatusCode: 0, StatusMsg: "success", Token: token, UserID: uid,
		})

	}
	ctx.JSON(http.StatusBadRequest, registeresponse{
		StatusCode: -1, StatusMsg: "fail", Token: "", UserID: 0,
	})

	tx.Commit()
}
