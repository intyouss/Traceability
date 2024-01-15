package dao

import (
	"errors"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
	"gorm.io/gorm"
)

var UserDaoIns *UserDao

type UserDao struct {
	*BaseDao
}

func NewUserDao() *UserDao {
	if UserDaoIns == nil {
		UserDaoIns = &UserDao{
			BaseDao: NewBaseDao(),
		}
	}
	return UserDaoIns
}

func (u *UserDao) GetUserByName(username string) (*models.User, error) {
	var user *models.User
	err := u.DB.Model(&user).Where("username = ?", username).First(&user).Error
	return user, err
}

// GetUserByNameAndPassword 根据用户名和密码获取用户
func (u *UserDao) GetUserByNameAndPassword(username, password string) (*models.User, error) {
	var user *models.User
	err := u.DB.Model(&models.User{}).Where("username = ? AND password = ?", username, password).First(&user).Error
	return user, err
}

// CheckUserNameExist 检查用户名是否存在
func (u *UserDao) CheckUserNameExist(username string) bool {
	var user models.User
	err := u.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

// AddUser 添加用户
func (u *UserDao) AddUser(userAddDTO *dto.UserAddDTO) (*models.User, error) {
	var user models.User
	userAddDTO.ToModel(&user)
	err := u.DB.Create(&user).Error
	if err == nil {
		userAddDTO.ID = user.ID
		userAddDTO.Password = ""
	}
	return &user, err
}

// GetUserById 根据id获取用户
func (u *UserDao) GetUserById(id uint) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&models.User{}).First(&user, id).Error
	return &user, err
}

// GetUserList 获取用户列表
func (u *UserDao) GetUserList(userListDto *dto.UserListDTO) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64
	err := u.DB.Model(&models.User{}).
		Scopes(Paginate(userListDto.CommonPageDTO)).Find(&users).
		Offset(-1).Limit(-1).Count(&total).Error
	return users, total, err
}

func (u *UserDao) UpdateUser(updateDTO *dto.UserUpdateDTO) error {
	var user models.User
	err := u.DB.First(&user, updateDTO.ID).Error
	if err != nil {
		return err
	}
	updateDTO.ToModel(&user)
	return u.DB.Save(&user).Error
}

func (u *UserDao) DeleteUserById(id uint) error {
	return u.DB.Delete(&models.User{}, id).Error
}
