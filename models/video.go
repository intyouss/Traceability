package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	// 视频标题
	Title string `gorm:"type:varchar(256);not null" json:"title"`
	// 视频地址
	PlayUrl string `gorm:"type:varchar(256);not null" json:"play_url"`
	// 视频封面地址
	CoverUrl string `gorm:"type:varchar(256);not null" json:"cover_url"`
	// 视频作者
	AuthorID uint `gorm:"not null" json:"author_id"`
	// 用户喜爱数
	LikeCount uint `gorm:"not null" json:"like_count"`
	// 用户评论数
	CommentCount uint `gorm:"not null" json:"comment_count"`
}
