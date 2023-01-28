package models

type TiktokMessage struct {
	// 消息id
	Mid int64 `json:"mid" form:"mid" gorm:"primaryKey" `
	// 消息内容
	Content string `json:"content" form:"content" `
	// 消息发送时间
	CreateTime string `json:"createTime" form:"createTime" gorm:"autoCreateTime" `
}
