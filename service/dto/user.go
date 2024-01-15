package dto

import "github.com/intyouss/Traceability/models"

type UserLoginDto struct {
	Username string `json:"username" form:"username" binding:"required" message:"Username cannot be empty"`
	Password string `json:"password" form:"password" binding:"required" message:"Password cannot be empty"`
}

// UserAddDTO 用户添加数据传输对象
type UserAddDTO struct {
	ID       uint
	Username string `json:"username" form:"username" binding:"required" message:"Username cannot be empty"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"Password cannot be empty"`
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
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
}

func (u *UserUpdateDTO) ToModel(user *models.User) {
	user.ID = u.ID
	user.Username = u.Username
	user.Avatar = u.Avatar
	user.Email = u.Email
	user.Mobile = u.Mobile
}

type UserListDTO struct {
	CommonPageDTO
}
