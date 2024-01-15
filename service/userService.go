package service

import (
	"errors"
	"fmt"
	"github.com/intyouss/Traceability/dao"
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
func (u *UserService) Login(userDTO dto.UserLoginDto) (*models.User, string, error) {
	var token string
	user, err := u.Dao.GetUserByName(userDTO.Username)
	if err != nil || !utils.ComparePassword(user.Password, userDTO.Password) {
		return nil, "", errors.New("invalid username or password")
	} else {
		// 生成token
		token, err = utils.GenerateToken(user.ID, user.Username)
		if err != nil {
			return nil, "", errors.New(fmt.Sprintf("generate token error: %s", err.Error()))
		}
	}
	return user, token, nil
}

// Register 用户注册
func (u *UserService) Register(userAddDTO *dto.UserAddDTO) (*models.User, string, error) {
	var token string
	if u.Dao.CheckUserNameExist(userAddDTO.Username) {
		return nil, "", errors.New("username already exists")
	}
	user, err := u.Dao.AddUser(userAddDTO)
	if err != nil {
		return nil, "", errors.New(fmt.Sprintf("add user error: %s", err.Error()))
	}
	// 生成token
	token, err = utils.GenerateToken(userAddDTO.ID, userAddDTO.Username)
	if err != nil {
		return nil, "", errors.New(fmt.Sprintf("generate token error: %s", err.Error()))
	}
	return user, token, nil
}

// GetUserById 根据id获取用户
func (u *UserService) GetUserById(idDTO *dto.CommonIDDTO) (*models.User, error) {
	return u.Dao.GetUserById(idDTO.ID)
}

// GetUserList 获取用户列表
func (u *UserService) GetUserList(userListDTO *dto.UserListDTO) ([]*models.User, int64, error) {
	return u.Dao.GetUserList(userListDTO)
}

// UpdateUser 更新用户信息
func (u *UserService) UpdateUser(updateDTO *dto.UserUpdateDTO) error {
	if updateDTO.ID == 0 {
		return errors.New("invalid user id")
	}
	return u.Dao.UpdateUser(updateDTO)
}

// DeleteUserById 删除用户
func (u *UserService) DeleteUserById(commonIDDTO *dto.CommonIDDTO) error {
	return u.Dao.DeleteUserById(commonIDDTO.ID)
}
