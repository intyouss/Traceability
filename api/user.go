package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
)

const (
	ErrCodeRegister = iota + 10001
	ErrCodeGetUserById
	ErrCodeGetUserList
	ErrCodeUpdateUser
	ErrCodeUploadAvatar
	ErrCodeDeleteUser
	ErrCodeAbolishAvatarUpload
	ErrCodeLogin
	ErrCodeGetUserIncrease
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// Login 登录Api
// @Summary 用户登录
// @Description 用户登录
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Failure 401 {string} Response
// @Router /api/v1/public/user/login [post]
func (u UserApi) Login(ctx *gin.Context) {
	// 绑定并验证参数
	var loginDto dto.UserLoginDto
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &loginDto}).GetError(); err != nil {
		u.Fail(&Response{Code: ErrCodeLogin, Msg: err.Error()})
		return
	}

	// 调用service
	user, token, err := u.Service.Login(ctx, loginDto)
	if err != nil {
		u.Fail(&Response{
			Code: ErrCodeLogin,
			Msg:  err.Error()})
		return
	}

	u.Success(&Response{
		Data: gin.H{
			"user":  user,
			"token": token,
		},
	})
}

// Register 注册Api
// @Summary 用户注册
// @Description 用户注册
// @Param username formData string true "用户名"
// @Param password formData string true "旧密码"
// @Param new_password formData string true "新密码"
// @Param email formData string false "邮箱"
// @Param mobile formData string false "手机号"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/public/user/register [post]
func (u UserApi) Register(ctx *gin.Context) {
	var userAddDto dto.UserAddDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userAddDto}).GetError(); err != nil {
		u.Fail(&Response{Code: ErrCodeRegister, Msg: err.Error()})
		return
	}
	user, token, err := u.Service.Register(ctx, &userAddDto)
	if err != nil {
		u.Fail(&Response{Code: ErrCodeRegister, Msg: err.Error()})
		return
	}

	u.Success(&Response{
		Data: gin.H{
			"user":  user,
			"token": token,
		},
	})
}

// GetUserInfo 获取其他用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Param user_id query int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/public/user/ [get]

// GetUserInfo 获取用户本人信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Param token header string true "token"
// @Param user_id query int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/user/ [get]
func (u UserApi) GetUserInfo(ctx *gin.Context) {
	var idDto dto.CommonIDDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &idDto}).GetError(); err != nil {
		u.Fail(&Response{Code: ErrCodeGetUserById, Msg: err.Error()})
		return
	}

	user, err := u.Service.GetUserById(ctx, &idDto)
	if err != nil {
		u.Fail(&Response{Code: ErrCodeGetUserById, Msg: err.Error()})
		return
	}

	u.Success(&Response{
		Data: gin.H{
			"user": user,
		},
	})
}

// GetUserList 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Param key query string false "关键字"
// @Param page query int false "页码"
// @Param limit query int false "每页数量"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/public/user/list [get]

// GetUserList 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Param token header string true "token"
// @Param key query string false "关键字"
// @Param page query int false "页码"
// @Param limit query int false "每页数量"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/user/list [get]
func (u UserApi) GetUserList(ctx *gin.Context) {
	var userListDto dto.UserListDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userListDto}).GetError(); err != nil {
		u.Fail(&Response{Code: ErrCodeGetUserList, Msg: err.Error()})
		return
	}

	users, total, err := u.Service.GetUserList(ctx, &userListDto)
	if err != nil {
		u.Fail(&Response{Code: ErrCodeGetUserList, Msg: err.Error()})
		return
	}
	if len(users) == 0 {
		u.Success(&Response{
			Data: gin.H{
				"users": []*dto.User{},
			},
		})
		return
	}

	u.Success(&Response{
		Data: gin.H{
			"users": users,
		},
		Total: total,
	})
}

// UpdateUser 更新用户信息
// @Summary 更新用户信息
// @Description 更新用户信息
// @Param token header string true "token"
// @Param password formData string false "密码"
// @Param email formData string false "邮箱"
// @Param mobile formData string false "手机号"
// @Param avatar formData string false "头像地址"
// @Param signature formData string false "个性签名"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/user/update [post]
func (u UserApi) UpdateUser(ctx *gin.Context) {
	var updateDTO dto.UserUpdateDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &updateDTO}).GetError(); err != nil {
		u.Fail(&Response{Code: ErrCodeUpdateUser, Msg: err.Error()})
		return
	}

	err := u.Service.UpdateUser(ctx, &updateDTO)
	if err != nil {
		u.Fail(&Response{Code: ErrCodeUpdateUser, Msg: err.Error()})
		return
	}

	u.Success(&Response{})
}

// UploadAvatar 上传头像
// @Summary 上传头像
// @Description 上传头像
// @Param token header string true "token"
// @Param avatar formData file true "头像"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/user/upload/avatar [post]
func (u UserApi) UploadAvatar(ctx *gin.Context) {
	var avatarDTO dto.UploadAvatarDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &avatarDTO}).GetError(); err != nil {
		u.Fail(&Response{Code: ErrCodeUploadAvatar, Msg: err.Error()})
		return
	}

	avatarUrl, err := u.Service.UploadAvatar(ctx, &avatarDTO)
	if err != nil {
		u.Fail(&Response{Code: ErrCodeUploadAvatar, Msg: err.Error()})
		return
	}

	u.Success(&Response{
		Data: gin.H{
			"avatar_url": avatarUrl,
		},
	})
}

// AbolishAvatarUpload 取消头像上传
// @Summary 取消头像上传
// @Description 取消头像上传
// @Param token header string true "token"
// @Param user_id formData int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/user/upload/avatar/abolish [post]
func (u UserApi) AbolishAvatarUpload(ctx *gin.Context) {
	var abolishDTO dto.AbolishAvatarDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &abolishDTO}).GetError(); err != nil {
		u.Fail(&Response{Code: ErrCodeAbolishAvatarUpload, Msg: err.Error()})
		return
	}

	err := u.Service.DeleteRemoteAvatar(ctx, &abolishDTO)
	if err != nil {
		u.Fail(&Response{Code: ErrCodeAbolishAvatarUpload, Msg: err.Error()})
		return
	}

	u.Success(&Response{})
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Param token header string true "token"
// @Param user_id formData int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/user/delete [post]
func (u UserApi) DeleteUser(ctx *gin.Context) {
	var idDTO dto.CommonIDDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &idDTO}).GetError(); err != nil {
		u.Fail(&Response{Code: ErrCodeDeleteUser, Msg: err.Error()})
		return
	}

	err := u.Service.DeleteUserById(ctx, &idDTO)
	if err != nil {
		u.Fail(&Response{Code: ErrCodeDeleteUser, Msg: err.Error()})
		return
	}
	u.Success(&Response{})
}

// GetUserIncrease 获取月总日用户增长记录列表
// @Summary 获取月总日用户增长记录列表
// @Description 获取月总日用户增长记录列表
// @Param token header string true "token"
// @Param year query string true "年份"
// @Param month query string true "月份"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/user/increase [get]
func (u UserApi) GetUserIncrease(ctx *gin.Context) {
	var list dto.UserIncreaseListDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &list}).GetError(); err != nil {
		u.Fail(&Response{Code: ErrCodeGetUserIncrease, Msg: err.Error()})
		return
	}

	c, err := u.Service.GetUserIncrease(ctx, &list)
	if err != nil {
		u.Fail(&Response{Code: ErrCodeGetUserIncrease, Msg: err.Error()})
		return
	}
	u.Success(&Response{
		Data: gin.H{
			"user_increase_list": c,
		},
	})
}
