package model

import (
	"time"
)

type User struct {
	// user_id
	UserId string `json:"userId" form:"userId" gorm:"primaryKey" `
	// 账号
	Useraccount string `json:"useraccount" form:"useraccount" `
	// 用户昵称
	Username string `json:"username" form:"username" `
	// 密码
	Userpassword string `json:"userpassword" form:"userpassword" `
	// 用户头像
	Useravatar string `json:"useravatar" form:"useravatar" `
	// 性别
	Gender int64 `json:"gender" form:"gender" `
	// 关注数
	follow_count int64 `json:"follow_count" form:"follow_count" `
	// 粉丝数
	follower_count int64 `json:"follower_count" form:"follower_count" `
	// 更新时间
	UpdateTime time.Time `json:"updateTime" form:"updateTime" `
	// 创建时间
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"autoCreateTime" `
	// 是否删除
	IsDelete int64 `json:"isDelete" form:"isDelete" `
}
