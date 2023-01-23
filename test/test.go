package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
