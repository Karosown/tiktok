package models

type Status struct {
	StatusCode int64  `gorm:"column:status_code" json:"StatusCode"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `gorm:"column:status_msg" json:"StatusMsg"`   // 返回状态描述
}

type Video struct {
	Vid           int64  `gorm:"column:vid" json:"Vid"`                      // 视频唯一标识
	Author        User   `gorm:"column:author" json:"Author"`                // 视频作者信息
	PlayUrl       string `gorm:"column:play_url" json:"PlayUrl"`             // 视频播放地址
	CoverUrl      string `gorm:"column:cover_url" json:"CoverUrl"`           // 视频封面地址
	FavoriteCount int64  `gorm:"column:favorite_count" json:"FavoriteCount"` // 视频的点赞总数
	CommentCount  int64  `gorm:"column:comment_count" json:"CommentCount"`   // 视频的评论总数
	IsFavorite    bool   `gorm:"column:is_favorite" json:"IsFavorite"`       // true-已点赞，false-未点赞
	Title         string `gorm:"column:title" json:"Title"`                  // 视频标题
}

type User struct {
	Uid           int64  `gorm:"column:uid" json:"Uid"`                      // 用户id
	Name          string `gorm:"column:name" json:"Name"`                    // 用户名称
	FollowCount   int64  `gorm:"column:follow_count" json:"FollowCount"`     // 关注总数
	FollowerCount int64  `gorm:"column:follower_count" json:"FollowerCount"` // 粉丝总数
	IsFollow      bool   `gorm:"column:is_follow" json:"IsFollow"`           // true-已关注，false-未关注
}

type Comment struct {
	Cid        int64  `gorm:"column:cid" json:"Cid"`                // 评论id
	User       User   `gorm:"column:user" json:"User"`              // 评论用户信息
	Content    string `gorm:"column:content" json:"Content"`        // 评论内容
	CreateDate string `gorm:"column:create_date" json:"CreateDate"` // 评论发布日期，格式 mm-dd
}

type Message struct {
	Mid        string `gorm:"column:mid" json:"Mid"`                // 消息id
	Content    string `gorm:"column:content" json:"Content"`        // 消息内容
	CreateTime int64  `gorm:"column:create_time" json:"CreateTime"` // 消息发送时间 yyyy-MM-dd HH:MM:ss
}

// UserAccount  string `grom:"userAccount"`  //账号
// UserName     string `grom:"UserName"`     //用户昵称
// UserPassword string `grom:"UserPassword"` //密码
// UserAvater   string `grom:"UserAvater"`   //头像
// Gender       int16  `grom:"Gender"`       //性别
// FollowCount  int64  `grom:"FollowCount"`  //关注数
// FlowerCount  int64  `grom:"FlowerCount"`  //粉丝数
// UpdateTime   string `grom:"UpdateTime"`   //更新时间
// CreateTime   string `grom:"CreateTime"`   //创建时间
// IsDelete     int16  `grom:"IsDelete"`     //是否删除
//omitempty字段优化
