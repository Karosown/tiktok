package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type loginmsg struct {
// 	Uid      int64
// 	Username string
// 	models.Status
// }

func Login(ctx *gin.Context) {
	ctx.Query("username")
	ctx.Query("password")
	//db := dao.GetDB(ctx)

	ctx.JSON(http.StatusOK, gin.H{})
}
