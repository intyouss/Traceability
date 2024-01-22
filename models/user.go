package models

import (
	"github.com/intyouss/Traceability/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"type:varchar(64);not null;unique" json:"username"`
	Password    string `gorm:"type:varchar(128);not null" json:"-"`
	Avatar      string `gorm:"type:varchar(256)" json:"avatar"`
	Email       string `gorm:"type:varchar(128);unique" json:"email"`
	Mobile      string `gorm:"type:varchar(11);unique" json:"mobile"`
	FollowCount uint   `gorm:"not null;default:0" json:"follow_count"`
	FansCount   uint   `gorm:"not null;default:0" json:"fans_count"`
	LikeCount   uint   `gorm:"not null;default:0" json:"like_count"`
	LikedCount  uint   `gorm:"not null;default:0" json:"liked_count"`
	VideoCount  uint   `gorm:"not null;default:0" json:"video_count"`
	Signature   string `gorm:"type:varchar(256)" json:"signature"`
}

// Encrypt 密码加密
func (u *User) Encrypt() error {
	hashPassword, err := utils.Encrypt(u.Password)
	if err == nil {
		u.Password = hashPassword
	}
	return err
}

// BeforeCreate BeforeHook钩子
func (u *User) BeforeCreate(db *gorm.DB) error {
	return u.Encrypt()
}

type LoginUser struct {
	ID       uint
	Username string
}
