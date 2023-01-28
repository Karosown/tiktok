package models

type TiktokUser struct {
	// 用户id
	Uid int64 `json:"uid" form:"uid" gorm:"primaryKey" `
	// 用户名称
	Name string `json:"name" form:"name" `
	// 用户账号
	Username string `json:"username" form:"username" `
	// 用户密码
	Userpassword string `json:"userpassword" form:"userpassword" `
	// 关注总数
	FollowCount interface{} `json:"followCount" form:"followCount" `
	// 粉丝总数
	FollowerCount interface{} `json:"followerCount" form:"followerCount" `
	// 是否关注
	IsFollow int64 `json:"isFollow" form:"isFollow" `
	// 用户授权token
	Token string `json:"token" form:"token" `
}
