package web

import (
	"fmt"
	"tiktok/logic"

	"github.com/gin-gonic/gin"
)

func RouterInit(router *gin.Engine) {
	apiRouter := router.Group("/douyin")
	{
		//basic api
		apiRouter.GET("/feed/", logic.Feed)
		// apiRouter.POST("/user/register/")
		// apiRouter.POST("/user/login/")
		// apiRouter.GET("/user/")
		// apiRouter.POST("/publish/action/")
		// apiRouter.GET("/publish/list/")

		// //interaction api
		// apiRouter.POST("/favorite/action/")
		// apiRouter.GET("/favorite/list/")
		// apiRouter.POST("/comment/action/")
		// apiRouter.GET("/comment/list/")

		// //relation api
		// apiRelation := apiRouter.Group("relation")
		// {
		// 	apiRelation.POST("/action/")
		// 	apiRelation.POST("/follow/list/")
		// 	apiRelation.POST("/follower/list/")
		// 	apiRelation.POST("/friend/list/")
		// }
	}
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("dsasf")
	}
}
