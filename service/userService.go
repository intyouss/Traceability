package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
	"github.com/intyouss/Traceability/utils"
)

var UserServiceIns *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if UserServiceIns == nil {
		UserServiceIns = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return UserServiceIns
}

// Login 用户登录
func (u *UserService) Login(ctx context.Context, userDTO dto.UserLoginDto) (*models.User, string, error) {
	var token string
	user, err := u.Dao.GetUserByName(ctx, userDTO.Username)
	if err != nil || !utils.ComparePassword(user.Password, userDTO.Password) {
		return nil, "", errors.New("invalid username or password")
	} else {
		// 生成token
		token, err = utils.GenerateToken(user.ID, user.Username)
		if err != nil {
			return nil, "", fmt.Errorf("generate token error: %s", err.Error())
		}
	}
	return user, token, nil
}

// Register 用户注册
func (u *UserService) Register(ctx context.Context, userAddDTO *dto.UserAddDTO) (*models.User, string, error) {
	var token string
	if u.Dao.CheckUserNameExist(userAddDTO.Username) {
		return nil, "", errors.New("username already exists")
	}
	user, err := u.Dao.AddUser(ctx, userAddDTO)
	if err != nil {
		return nil, "", fmt.Errorf("add user error: %s", err.Error())
	}
	// 生成token
	token, err = utils.GenerateToken(userAddDTO.ID, userAddDTO.Username)
	if err != nil {
		return nil, "", fmt.Errorf("generate token error: %s", err.Error())
	}
	return user, token, nil
}

// GetUserById 根据id获取用户
func (u *UserService) GetUserById(ctx context.Context, idDTO *dto.CommonIDDTO) (*models.User, error) {
	return u.Dao.GetUserById(ctx, idDTO.ID)
}

// GetUserListBySearch 模糊搜索用户列表
func (u *UserService) GetUserListBySearch(
	ctx context.Context, userListDTO *dto.UserSearchListDTO,
) ([]*models.User, int64, error) {
	return u.Dao.GetUserListBySearch(ctx, userListDTO)
}

// UpdateUser 更新用户信息
func (u *UserService) UpdateUser(ctx context.Context, updateDTO *dto.UserUpdateDTO) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	if updateDTO.ID != userId {
		return errors.New("don't have permission")
	}
	return u.Dao.UpdateUser(ctx, updateDTO)
}

// DeleteUserById 删除用户
func (u *UserService) DeleteUserById(ctx context.Context, idDTO *dto.CommonIDDTO) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	if idDTO.ID != userId {
		return errors.New("don't have permission")
	}
	return u.Dao.DeleteUserById(ctx, userId)
}

// GetUserListByIds 根据id列表获取用户列表
func (u *UserService) GetUserListByIds(ctx context.Context, ids []uint) ([]*models.User, error) {
	return u.Dao.GetUserListByIds(ctx, ids)
}

// IsExist 用户是否存在
func (u *UserService) IsExist(ctx context.Context, id uint) bool {
	return u.Dao.IsExist(ctx, id)
}
