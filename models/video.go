package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	// 视频标题
	Title string `gorm:"type:varchar(256);not null;unique" json:"title"`
	// 视频地址
	PlayUrl string `gorm:"type:text;not null" json:"play_url"`
	// 视频封面地址
	CoverUrl string `gorm:"type:text;not null" json:"cover_url"`
	// 视频作者
	AuthorID uint `gorm:"not null" json:"author_id"`
	// 用户喜爱数
	LikeCount uint `gorm:"not null" json:"like_count"`
	// 用户评论数
	CommentCount uint `gorm:"not null" json:"comment_count"`
	// 用户收藏数
	CollectCount uint `gorm:"not null" json:"collect_count"`
}

type VideoIncrease struct {
	gorm.Model
	Year  uint `gorm:"not null;" json:"year"`
	Month uint `gorm:"not null;" json:"month"`
	Day   uint `gorm:"not null;" json:"day"`
	Count uint `gorm:"not null;default:0" json:"count"`
}

func (v *VideoIncrease) SetDate() {
	v.Year = uint(v.CreatedAt.Year())
	v.Month = uint(v.CreatedAt.Month())
	v.Day = uint(v.CreatedAt.Day())
}

// BeforeCreate BeforeHook钩子
func (v *VideoIncrease) BeforeCreate(db *gorm.DB) error {
	v.SetDate()
	return nil
}
