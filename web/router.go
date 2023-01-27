package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"os"
	"tiktok/dao"
	"tiktok/logic"
)

var DB gorm.DB

func InitServer(router *gin.Engine) {
	//初始化路由
	RouterInit(router)

	//初始化日志
	LogInit()

	//初始化数据库
	DB = *DBInit()
}

func RouterInit(router *gin.Engine) {
	apiRouter := router.Group("/douyin")
	{
		//basic api
		apiRouter.GET("/feed/", logic.Feed) //中间件和路由在同一层时会优先执行路由
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

func Loginit() {
	logFile, err := os.OpenFile("./web/log/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func DBInit() *gorm.DB {
	return dao.DBConfig()
}

func GetDB() *gorm.DB {
	return &DB
}
