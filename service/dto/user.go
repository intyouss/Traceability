package dto

import "github.com/intyouss/Traceability/models"

// User 用户数据传输对象
type User struct {
	ID uint `json:"id"`
	// 密码
	Password string `json:"-"`
	// 用户名
	Username string `json:"username"`
	// 头像
	Avatar string `json:"avatar"`
	// 邮箱
	Email string `json:"email"`
	// 手机号
	Mobile string `json:"mobile"`
	// 关注数
	FollowCount uint `json:"follow_count"`
	// 粉丝数
	FansCount uint `json:"fans_count"`
	// 视频点赞数
	LikeCount uint `json:"like_count"`
	// 视频被点赞数
	LikedCount uint `json:"liked_count"`
	// 视频数
	VideoCount uint `json:"video_count"`
	// 个性签名
	Signature string `json:"signature"`
}

type UserLoginDto struct {
	Username string `json:"username" form:"username" binding:"required" message:"username cannot be empty"`
	Password string `json:"password" form:"password" binding:"required" message:"password cannot be empty"`
}

// UserAddDTO 用户添加数据传输对象
type UserAddDTO struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username" binding:"required" message:"username cannot be empty"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"password cannot be empty"`
	Avatar   string
	Email    string `json:"email" form:"email"`
	Mobile   string `json:"mobile" form:"mobile"`
}

func (u *UserAddDTO) ToModel(user *models.User) {
	user.Username = u.Username
	user.Email = u.Email
	user.Mobile = u.Mobile
	user.Password = u.Password
}

type UserUpdateDTO struct {
	UserID   uint   `json:"user_id" form:"user_id" binding:"required" message:"user_id cannot be empty"`
	Username string `json:"username" form:"username"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
}

func (u *UserUpdateDTO) ToModel(user *models.User) {
	user.ID = u.UserID
	user.Username = u.Username
	user.Avatar = u.Avatar
	user.Email = u.Email
	user.Mobile = u.Mobile
}

type UserListDTO struct {
	CommonPageDTO
}
