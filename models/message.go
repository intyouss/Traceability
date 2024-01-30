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
