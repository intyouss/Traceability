package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	// 源用户id
	FromUserID uint `gorm:"not null" json:"from_user_id"`
	// 目标用户id
	ToUserID uint `gorm:"not null" json:"to_user_id"`
	// 消息内容
	Content string `gorm:"not null" json:"content"`
}

// MessageOpen 消息打开表
type MessageOpen struct {
	gorm.Model
	// 用户id
	UserID uint `gorm:"not null" json:"user_id"`
	// 被打开用户id
	OpenUserID uint `gorm:"not null" json:"open_user_id"`
	// 正在使用的人数
	UseCount uint `gorm:"not null;default:1" json:"use_count"`
}
