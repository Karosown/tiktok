package dbtest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type People struct {
	Name string `json:"name"`
}

func Test(ctx *gin.Context) {
	obj := People{
		Name: "zhangsan",
	}

	ctx.JSON(http.StatusOK, obj)
}
