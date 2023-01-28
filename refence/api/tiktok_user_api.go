package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func init() {
	c := &TiktokUserApi{}
	c.init(app.R) //这里需要引入你的gin框架的实例
}

func (t TiktokUserApi) init(g *gin.Engine) {
	// 依次: 分页列表，单条，插入，修改，删除
	group := g.Group("tiktok/user")
	group.GET("/list", t.list) //不设置限制条件的画默认查询所有
	group.GET("/one", t.one)
	group.POST("/insert", t.insert)
	group.POST("/update", t.update)
	group.POST("/delete", t.delete)
}

// TiktokUserApi 控制器
type TiktokUserApi struct {
	Page int   `form:"page"`
	Size int   `form:"size"`
	Ids  []int `form:"ids"`
}

func (t TiktokUserApi) db() *gorm.DB {
	return common.Db
}

// 分页列表
func (t TiktokUserApi) list(c *gin.Context) {
	_ = c.Bind(&t)
	v := models.TiktokUser{}
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, model.OkData(service.TiktokUserService.List(t.Page, t.Size, &v)))
}

// 根据主键Id查询记录
func (t TiktokUserApi) one(c *gin.Context) {
	var v models.TiktokUser
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, model.OkData(service.TiktokUserService.One(v.Uid)))
}

// 修改记录
func (t TiktokUserApi) update(c *gin.Context) {
	var v models.TiktokUser
	_ = c.ShouldBindJSON(&v)
	service.TiktokUserService.Update(v)
	c.JSON(http.StatusOK, model.OkMsg("修改成功！"))
}

// 插入记录
func (t TiktokUserApi) insert(c *gin.Context) {
	var v models.TiktokUser
	_ = c.ShouldBindJSON(&v)
	service.TiktokUserService.Insert(v)
	c.JSON(http.StatusOK, model.OkMsg("插入成功！"))
}

// 根据主键删除
func (t TiktokUserApi) delete(c *gin.Context) {
	_ = c.Bind(&t)
	service.TiktokUserService.Delete(t.Ids)
	c.JSON(http.StatusOK, model.OkMsg("删除成功！"))
}
