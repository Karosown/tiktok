package web

import (
	"fmt"
	"log"
	"os"
	"tiktok/logic"

	"github.com/gin-gonic/gin"
)

func InitServer(router *gin.Engine) {
	//初始化路由
	RouterInit(router)

	//初始化数据库
}

func RouterInit(router *gin.Engine) {
	apiRouter := router.Group("/douyin")
	{
		//basic api
		apiRouter.GET("/feed/", CheckToken(), logic.Feed) //中间件和路由在同一层时会优先执行路由
		// apiRouter.POST("/user/register/")
		apiRouter.POST("/user/login/", logic.Login)
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

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//验证token
	}
}

func init() {
	logFile, err := os.OpenFile("./web/log/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
