package models

type TiktokVideo struct {
	// 视频唯一标识
	Vid int64 `json:"vid" form:"vid" gorm:"primaryKey" `
	// 视频作者id
	Author int64 `json:"author" form:"author" `
	// 视频播放地址
	PlayUrl string `json:"playUrl" form:"playUrl" `
	// 视频封面地址
	CoverUrl string `json:"coverUrl" form:"coverUrl" `
	// 视频的点赞总数
	FavoriteCount interface{} `json:"favoriteCount" form:"favoriteCount" `
	// 视频的评论总数
	CommentCount interface{} `json:"commentCount" form:"commentCount" `
	// 是否点赞
	IsFavorite int64 `json:"isFavorite" form:"isFavorite" `
	// 视频标题
	Title string `json:"title" form:"title" `
	// 时间
	NextTime int64 `json:"nextTime" form:"nextTime" `
}
