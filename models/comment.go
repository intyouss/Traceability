package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:varchar(256);not null" json:"content"`
	// 评论人
	UserId uint `gorm:"not null" json:"user_id"`
	// 评论的视频
	VideoId uint `gorm:"not null" json:"video_id"`
}
