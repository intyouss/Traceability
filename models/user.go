package models

import (
	"github.com/intyouss/Traceability/utils"
	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model
	// 用户名
	Username string `gorm:"type:varchar(64);not null;unique" json:"username"`
	// 密码
	Password string `gorm:"type:varchar(128);not null" json:"-"`
	// 头像
	Avatar string `gorm:"type:text" json:"avatar"`
	// 邮箱
	Email string `gorm:"type:varchar(128)" json:"email"`
	// 手机号
	Mobile string `gorm:"type:varchar(11)" json:"mobile"`
	// 关注数
	FocusCount uint `gorm:"not null;default:0" json:"focus_count"`
	// 粉丝数
	FansCount uint `gorm:"not null;default:0" json:"fans_count"`
	// 点赞数
	LikeCount uint `gorm:"not null;default:0" json:"like_count"`
	// 被点赞数
	LikedCount uint `gorm:"not null;default:0" json:"liked_count"`
	// 视频数
	VideoCount uint `gorm:"not null;default:0" json:"video_count"`
	// 收藏数
	CollectCount uint `gorm:"not null;default:0" json:"collect_count"`
	// 个性签名
	Signature string `gorm:"type:varchar(256)" json:"signature"`
	// 角色
	Role uint `gorm:"not null;default:1" json:"role"`
	// 状态
	Status uint `gorm:"not null;default:1" json:"status"`
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

// UserIncrease 用户增长
type UserIncrease struct {
	gorm.Model
	Year  uint `gorm:"not null;" json:"year"`
	Month uint `gorm:"not null;" json:"month"`
	Day   uint `gorm:"not null;" json:"day"`
	Count uint `gorm:"not null;default:1" json:"count"`
}

func (u *UserIncrease) SetDate() {
	u.Year = uint(u.CreatedAt.Year())
	u.Month = uint(u.CreatedAt.Month())
	u.Day = uint(u.CreatedAt.Day())
}

// BeforeCreate BeforeHook钩子
func (u *UserIncrease) BeforeCreate(db *gorm.DB) error {
	u.SetDate()
	return nil
}

// Role 角色
type Role struct {
	gorm.Model
	// 角色名
	Name string `gorm:"type:varchar(64);not null;unique" json:"name"`
	// 角色描述
	Desc string `gorm:"type:varchar(256)" json:"desc"`
	// 状态
	Status uint `gorm:"not null;default:1" json:"status"`
}

type LoginUser struct {
	ID       uint
	Username string
	Role     uint
}
