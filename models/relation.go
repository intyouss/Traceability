package models

import "gorm.io/gorm"

type Relation struct {
	gorm.Model
	// 用户id
	UserID uint `gorm:"not null" json:"user_id"`
	// 关注id
	FocusID uint `gorm:"not null" json:"focus_id"`
}
