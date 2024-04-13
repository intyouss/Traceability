package dto

import (
	"mime/multipart"

	"github.com/intyouss/Traceability/models"
)

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
	FocusCount uint `json:"focus_count"`
	// 粉丝数
	FansCount uint `json:"fans_count"`
	// 视频点赞数
	LikeCount uint `json:"like_count"`
	// 视频被点赞数
	LikedCount uint `json:"liked_count"`
	// 视频数
	VideoCount uint `json:"video_count"`
	// 收藏数
	CollectCount uint `json:"collect_count"`
	// 个性签名
	Signature string `json:"signature"`
	// 是否关注
	IsFocus bool `json:"is_focus"`
	// 角色
	Role *Role `json:"role"`
	// 状态
	Status uint `json:"status"`
}

type UserIncrease struct {
	Year  uint `json:"year"`
	Month uint `json:"month"`
	Day   uint `json:"day"`
	Count uint `json:"count"`
}

// Role 角色
type Role struct {
	ID uint `json:"id"`
	// 角色名
	Name string `json:"name"`
	// 角色描述
	Desc string `json:"desc"`
	// 状态
	Status uint `json:"status"`
}

type UserLoginDto struct {
	Username string `json:"username" form:"username" binding:"required" message:"username cannot be empty"`
	Password string `json:"password" form:"password" binding:"required" message:"password cannot be empty"`
	Admin    bool   `json:"admin" form:"admin"`
}

// UserAddDTO 用户添加数据传输对象
type UserAddDTO struct {
	ID       *uint  `json:"id" form:"id"`
	Username string `json:"username" form:"username" binding:"required" message:"username cannot be empty"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"password cannot be empty"`
	//Avatar   string `json:"avatar" form:"avatar"`
	Email  string `json:"email" form:"email"`
	Mobile string `json:"mobile" form:"mobile"`
}

func (u *UserAddDTO) ToModel(user *models.User) {
	user.Username = u.Username
	user.Email = u.Email
	user.Mobile = u.Mobile
	user.Password = u.Password
}

type UserUpdateDTO struct {
	UserID      uint   `json:"user_id" form:"user_id"`
	Password    string `json:"password" form:"password"`
	NewPassword string `json:"new_password" form:"new_password"`
	AvatarUrl   string `json:"avatar_url" form:"avatar_url"`
	Signature   string `json:"signature" form:"signature"`
	Mobile      string `json:"mobile" form:"mobile"`
	Email       string `json:"email" form:"email"`
	Role        uint   `json:"role" form:"role"`
	Status      uint   `json:"status" form:"status"`
}

type UploadAvatarDTO struct {
	AvatarData multipart.FileHeader `json:"avatar_data" form:"avatar_data" type:"blob" binding:"required" message:"avatar_data cannot be empty"`
}

type AbolishAvatarDTO struct {
	UserId uint `json:"user_id" form:"user_id" binding:"required" message:"user_id cannot be empty"`
}

func (u *UserUpdateDTO) ToModel(user *models.User) {
	user.Password = u.NewPassword
	user.Avatar = u.AvatarUrl
	user.Signature = u.Signature
	user.Email = u.Email
	user.Mobile = u.Mobile
	user.Role = u.Role
	user.Status = u.Status
}

type UserListDTO struct {
	Key string `json:"key" form:"key"`
	CommonPageDTO
}

type UserIncreaseListDTO struct {
	Year  uint `json:"year" form:"year" binding:"required" message:"year cannot be empty"`
	Month uint `json:"month" form:"month" binding:"required" message:"month cannot be empty"`
}

type RoleListDTO struct {
	Key string `json:"key" form:"key"`
	CommonPageDTO
}

type RoleAddDTO struct {
	Name string `json:"name" form:"name" binding:"required" message:"name cannot be empty"`
	Desc string `json:"desc" form:"desc" binding:"required" message:"desc cannot be empty"`
}

type RoleUpdateDTO struct {
	CommonIDDTO
	Name   string `json:"name" form:"name"`
	Desc   string `json:"desc" form:"desc"`
	Status uint   `json:"status" form:"status"`
}

type RoleDeleteDTO struct {
	IDs []uint `json:"ids" form:"ids" binding:"required" message:"ids cannot be empty"`
}
