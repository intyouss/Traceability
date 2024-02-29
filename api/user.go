package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
	"github.com/jinzhu/copier"
)

const (
	ErrCodeRegister = iota + 10001
	ErrCodeGetUserById
	ErrCodeGetUserList
	ErrCodeUpdateUser
	ErrCodeDeleteUser
	ErrCodeLogin
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
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeLogin, Msg: err.Error()})
		return
	}

	// 调用service
	userDao, token, err := u.Service.Login(ctx, loginDto)
	if err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{
			Code: ErrCodeLogin,
			Msg:  err.Error()})
		return
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)

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
// @Param password formData string true "密码"
// @Param email formData string false "邮箱"
// @Param mobile formData string false "手机号"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/public/user/register [post]
func (u UserApi) Register(ctx *gin.Context) {
	var userAddDto dto.UserAddDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userAddDto}).GetError(); err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeRegister, Msg: err.Error()})
		return
	}
	userDao, token, err := u.Service.Register(ctx, &userAddDto)
	if err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeRegister, Msg: err.Error()})
		return
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)

	u.Success(&Response{
		Data: gin.H{
			"user":  user,
			"token": token,
		},
	})
}

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Param token header string true "token"
// @Param user_id query int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/user/ [get]
// @Router /api/v1/public/user/ [get]
func (u UserApi) GetUserInfo(ctx *gin.Context) {
	var idDto dto.CommonUserIDDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &idDto}).GetError(); err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeGetUserById, Msg: err.Error()})
		return
	}

	userDao, err := u.Service.GetUserById(ctx, &idDto)
	if err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeGetUserById, Msg: err.Error()})
		return
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)

	u.Success(&Response{
		Data: gin.H{
			"user": user,
		},
	})
}

// GetUserListBySearch 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Param key formData string false "关键字"
// @Param page formData int false "页码"
// @Param limit formData int false "每页数量"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/public/user/search [get]
func (u UserApi) GetUserListBySearch(ctx *gin.Context) {
	var userListDto dto.UserSearchListDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userListDto}).GetError(); err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeGetUserList, Msg: err.Error()})
		return
	}

	usersDao, total, err := u.Service.GetUserListBySearch(ctx, &userListDto)
	if err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeGetUserList, Msg: err.Error()})
		return
	}
	var users = make([]*dto.User, len(usersDao))
	_ = copier.Copy(&users, &usersDao)

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
// @Param username formData string false "用户名"
// @Param password formData string false "密码"
// @Param email formData string false "邮箱"
// @Param mobile formData string false "手机号"
// @Param avatar formData file false "头像"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/user/update [post]
func (u UserApi) UpdateUser(ctx *gin.Context) {
	var updateDTO dto.UserUpdateDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &updateDTO}).GetError(); err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeUpdateUser, Msg: err.Error()})
		return
	}

	// 上传头像
	//file, err := ctx.FormFile("avatar")
	//if err == nil {
	//	filePath := fmt.Sprintf("./upload/%s", file.Filename)
	//	err = ctx.SaveUploadedFile(file, filePath)
	//	if err != nil {
	//		u.Fail(&Response{Code: ErrCodeUpdateUser, Msg: err.Error()})
	//	}
	//	updateDTO.Avatar = filePath
	//}

	err := u.Service.UpdateUser(ctx, &updateDTO)
	if err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeUpdateUser, Msg: err.Error()})
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
	var idDTO dto.CommonUserIDDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &idDTO}).GetError(); err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeDeleteUser, Msg: err.Error()})
		return
	}

	err := u.Service.DeleteUserById(ctx, &idDTO)
	if err != nil {
		u.Logger.Error(err)
		u.Fail(&Response{Code: ErrCodeDeleteUser, Msg: err.Error()})
		return
	}
	u.Success(&Response{})
}
