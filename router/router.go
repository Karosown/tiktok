package router

import (
	"tiktok/dao"
	"tiktok/logic"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DB gorm.DB

func InitServer(router *gin.Engine) {
	//初始化路由
	RouterInit(router)

	//初始化日志
	LogInit()

	//初始化数据库
	dao.DBInit()
}

func RouterInit(router *gin.Engine) {
	apiRouter := router.Group("/douyin")
	{
		//basic api
		apiRouter.GET("/feed/", logic.Feed) //中间件和路由在同一层时会优先执行路由
		apiRouter.POST("/user/register/", logic.Register)
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

func LogInit() {
	//定位文件夹
	path := "./web/log/test.log"

	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	 WithMaxAge 和 WithRotationCount二者只能设置一个
	`WithMaxAge` 设置文件清理前的最长保存时间
	`WithRotationCount` 设置文件清理前最多保存的个数
	*/
	// 下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，多余的自动清理掉。
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H",
		//创建新log与旧log位于同一目录下
		rotatelogs.WithLinkName(path),
		//两小时建一个新log文件
		rotatelogs.WithRotationTime(time.Duration(2)*time.Hour),
		//至多有三个log文件
		rotatelogs.WithRotationTime(3),
	)
	logrus.SetOutput(writer)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}) //log.SetFormatter(&log.JSONFormatter{})
}
