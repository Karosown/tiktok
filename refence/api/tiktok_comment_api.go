package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func init() {
	c := &TiktokCommentApi{}
	c.init(app.R) //这里需要引入你的gin框架的实例
}

func (t TiktokCommentApi) init(g *gin.Engine) {
	// 依次: 分页列表，单条，插入，修改，删除
	group := g.Group("tiktok/comment")
	group.GET("/list", t.list) //不设置限制条件的画默认查询所有
	group.GET("/one", t.one)
	group.POST("/insert", t.insert)
	group.POST("/update", t.update)
	group.POST("/delete", t.delete)
}

// TiktokCommentApi 控制器
type TiktokCommentApi struct {
	Page int   `form:"page"`
	Size int   `form:"size"`
	Ids  []int `form:"ids"`
}

func (t TiktokCommentApi) db() *gorm.DB {
	return common.Db
}

// 分页列表
func (t TiktokCommentApi) list(c *gin.Context) {
	_ = c.Bind(&t)
	v := models.TiktokComment{}
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, model.OkData(service.TiktokCommentService.List(t.Page, t.Size, &v)))
}

// 根据主键Id查询记录
func (t TiktokCommentApi) one(c *gin.Context) {
	var v models.TiktokComment
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, model.OkData(service.TiktokCommentService.One(v.Cid)))
}

// 修改记录
func (t TiktokCommentApi) update(c *gin.Context) {
	var v models.TiktokComment
	_ = c.ShouldBindJSON(&v)
	service.TiktokCommentService.Update(v)
	c.JSON(http.StatusOK, model.OkMsg("修改成功！"))
}

// 插入记录
func (t TiktokCommentApi) insert(c *gin.Context) {
	var v models.TiktokComment
	_ = c.ShouldBindJSON(&v)
	service.TiktokCommentService.Insert(v)
	c.JSON(http.StatusOK, model.OkMsg("插入成功！"))
}

// 根据主键删除
func (t TiktokCommentApi) delete(c *gin.Context) {
	_ = c.Bind(&t)
	service.TiktokCommentService.Delete(t.Ids)
	c.JSON(http.StatusOK, model.OkMsg("删除成功！"))
}
