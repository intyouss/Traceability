package dto

import "mime/multipart"

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
	// 创建时间
	CreatedAt string `json:"created_at"`
}

type VideoListDTO struct {
	Type       uint   `json:"type" form:"type" binding:"required" message:"type cannot be empty"`
	LatestTime *int64 `json:"latest_time" form:"latest_time" binding:"required" message:"latest_time cannot be empty"`
}

type VideoPublishDTO struct {
	Title          string               `json:"title" form:"title" binding:"required" message:"title cannot be empty"`
	CoverImageData multipart.FileHeader `json:"cover_image_data" form:"cover_image_data" type:"blob" binding:"required"  message:"cover_image_data cannot be empty"`
	Data           multipart.FileHeader `json:"data" form:"data" type:"blob" binding:"required" message:"data cannot be empty" `
}

type VideoDeleteDTO struct {
	VideoID int `json:"video_id" form:"video_id" binding:"required" message:"video_id cannot be empty"`
}

type VideoSearchDTO struct {
	Key  string `json:"key" form:"key" binding:"required" message:"key cannot be empty"`
	Type uint   `json:"type" form:"type" binding:"required" message:"type cannot be empty"`
}
