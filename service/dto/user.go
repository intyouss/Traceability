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
}

type UserLoginDto struct {
	Username string `json:"username" form:"username" binding:"required" message:"username cannot be empty"`
	Password string `json:"password" form:"password" binding:"required" message:"password cannot be empty"`
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
	Password    string `json:"password" form:"password"`
	NewPassword string `json:"new_password" form:"new_password"`
	AvatarUrl   string `json:"avatar_url" form:"avatar_url"`
	Signature   string `json:"signature" form:"signature"`
	Mobile      string `json:"mobile" form:"mobile"`
	Email       string `json:"email" form:"email"`
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
}

type UserSearchListDTO struct {
	Key string `json:"key" form:"key" binding:"required" message:"key cannot be empty"`
	CommonPageDTO
}
