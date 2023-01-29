package logic

import (
	"net/http"
	"tiktok/dao"
	"tiktok/models"
	"tiktok/web"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type loginmsg struct {
	Uid   int64  `json:"user_id"`
	Token string `json:"token"`
	models.Status
}

type dbres struct {
	Uid      int64  `gorm:"column:uid"`
	Password string `gorm:"column:password"`
}

func Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	db := dao.GetDB(ctx)
	user := dbres{} //make优化

	err := db.Table("tiktok_user").Select("uid, password").Where("name = ?", username).Find(&user).Error
	if err != nil || user.Password != password { //用户名不存在或密码错误
		ctx.JSON(http.StatusBadRequest, loginmsg{
			Status: models.Status{
				StatusCode: -1,
				StatusMsg:  "账号或密码错误！",
			},
		})
		return
	}

	restoken, err := web.CreateToken(user.Uid, username)
	if err != nil { //创建token失败
		logrus.Error("login,web.CreateToken:创建token失败...")
		ctx.JSON(http.StatusBadRequest, loginmsg{
			Status: models.Status{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误！",
			},
		})
		return
	}

	err = db.Table("tiktok_user").Where("name = ?", username).Update("token", restoken).Error
	if err != nil {
		logrus.Error("login,db.Updata:更新token失败...")
		ctx.JSON(http.StatusBadRequest, loginmsg{
			Status: models.Status{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误！",
			},
		})
		return
	}

	res := loginmsg{
		Uid:   user.Uid,
		Token: restoken,
		Status: models.Status{
			StatusCode: 0,
			StatusMsg:  "登录成功！",
		},
	}
	ctx.JSON(http.StatusOK, res)
}
