package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"

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
func (u *UserService) Login(ctx context.Context, userDTO dto.UserLoginDto) (*dto.User, string, error) {
	var token string
	userDao, err := u.Dao.GetUserByName(ctx, userDTO.Username)
	if err != nil || !utils.ComparePassword(userDao.Password, userDTO.Password) {
		return nil, "", errors.New("invalid username or password")
	} else {
		// 生成token
		token, err = utils.GenerateToken(userDao.ID, userDao.Username)
		if err != nil {
			return nil, "", fmt.Errorf("generate token error: %s", err.Error())
		}
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)
	return user, token, nil
}

// Register 用户注册
func (u *UserService) Register(ctx context.Context, userAddDTO *dto.UserAddDTO) (*dto.User, string, error) {
	var token string
	if u.Dao.CheckUserNameExist(userAddDTO.Username) {
		return nil, "", errors.New("username already exists")
	}
	userDao, err := u.Dao.AddUser(ctx, userAddDTO)
	if err != nil {
		return nil, "", fmt.Errorf("add user error: %s", err.Error())
	}
	// 生成token
	token, err = utils.GenerateToken(*userAddDTO.ID, userAddDTO.Username)
	if err != nil {
		return nil, "", fmt.Errorf("generate token error: %s", err.Error())
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)
	return user, token, nil
}

// GetUserById 根据id获取用户
func (u *UserService) GetUserById(ctx context.Context, idDTO *dto.CommonIDDTO) (*dto.User, error) {
	var userId = *idDTO.ID
	if *idDTO.ID == 0 {
		userId = ctx.Value(global.LoginUser).(models.LoginUser).ID
	}
	userDao, err := u.Dao.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)
	return user, nil
}

// GetUserListBySearch 模糊搜索用户列表
func (u *UserService) GetUserListBySearch(
	ctx context.Context, userListDTO *dto.UserSearchListDTO,
) ([]*dto.User, int64, error) {
	usersDao, total, err := u.Dao.GetUserListBySearch(ctx, userListDTO)
	if err != nil {
		return nil, 0, err
	}
	var users = make([]*dto.User, len(usersDao))
	_ = copier.Copy(&users, &usersDao)
	return users, total, nil
}

// UpdateUser 更新用户信息
func (u *UserService) UpdateUser(ctx context.Context, updateDTO *dto.UserUpdateDTO) error {
	return u.Dao.UpdateUser(ctx, updateDTO)
}

// DeleteUserById 删除用户
func (u *UserService) DeleteUserById(ctx context.Context, idDTO *dto.CommonIDDTO) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	if *idDTO.ID != userId {
		return errors.New("don't have permission")
	}
	return u.Dao.DeleteUserById(ctx, userId)
}
