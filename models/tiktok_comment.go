package models

type TiktokComment struct {
	// 评论id
	Cid int64 `json:"cid" form:"cid" gorm:"primaryKey" `
	// 评论用户id
	Uid int64 `json:"uid" form:"uid" `
	// 评论内容
	Content string `json:"content" form:"content" `
	// 评论发布日期
	CreateDate string `json:"createDate" form:"createDate" `
}
