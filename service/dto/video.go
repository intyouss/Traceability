package dto

// Video 视频数据传输对象
type Video struct {
	ID uint `json:"id"`
	// 视频作者
	Author *User `json:"author"`
	// 视频标题
	Title string `json:"title"`
	// 视频描述
	Description string `json:"description"`
	// 视频地址
	PlayUrl string `json:"play_url"`
	// 视频封面地址
	CoverUrl string `json:"cover_url"`
	// 用户喜爱数
	LikeCount uint `json:"like_count"`
	// 用户评论数
	CommentCount uint `json:"comment_count"`
	// 用户喜爱状态
	LikeStatus bool `json:"like_status"`
}

type VideoListDTO struct {
	CommonPageDTO
	LatestTime string `json:"latest_time" form:"latest_time"`
}

type VideoPublishDTO struct {
	UserID uint   `json:"id" form:"id" binding:"required" message:"UserId cannot be empty"`
	Title  string `json:"title" form:"title" binding:"required" message:"Title cannot be empty"`
	Data   []byte `json:"data" form:"data" binding:"required" message:"Data cannot be empty"`
}
