package models

type Status struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type Video struct {
	Vid           int64  `json:"id"`             // 视频唯一标识
	Author        User   `json:"author"`         // 视频作者信息
	PlayUrl       string `json:"play_url"`       // 视频播放地址
	CoverUrl      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string `json:"title"`          // 视频标题
	NextTime      int64  `json:"next_time"`      //时间
}

type User struct {
	Uid           int64  `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Token         string `json:"token"`          //用户授权Token
}

type Comment struct {
	Cid        int64  `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户信息
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
}

type Message struct {
	Mid        string `json:"id"`           // 消息id
	Content    string `json:"content"`      // 消息内容
	CreateTime int64  `json:" create_time"` // 消息发送时间 yyyy-MM-dd HH:MM:ss
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
