package models

import "gorm.io/gorm"

type Collect struct {
	gorm.Model
	// 用户id
	UserID uint `gorm:"not null" json:"user_id"`
	// 视频id
	VideoID uint `gorm:"not null" json:"video_id"`
}
