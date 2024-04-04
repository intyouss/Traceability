package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/copier"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
	"github.com/intyouss/Traceability/utils"
)

var UserServiceIns *UserService

type UserService struct {
	*BaseService
	Dao         *dao.UserDao
	RelationDao *dao.RelationDao
}

func NewUserService() *UserService {
	if UserServiceIns == nil {
		UserServiceIns = &UserService{
			BaseService: NewBaseService(),
			Dao:         dao.NewUserDao(),
			RelationDao: dao.NewRelationDao(),
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

	year := uint(time.Now().Year())
	month := uint(time.Now().Month())
	day := uint(time.Now().Day())

	// 更新用户增长记录
	ok, _, err := u.Dao.GetUserIncrease(ctx, year, month, day)
	if err != nil {
		return nil, "", fmt.Errorf("update user increase count error: %s", err.Error())
	}
	if ok {
		err = u.Dao.UpdateUserIncreaseCount(ctx, year, month, day, 1)
		if err != nil {
			return nil, "", fmt.Errorf("update user increase count error: %s", err.Error())
		}
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
	// 如果id为0，则获取当前登录用户
	if *idDTO.ID == 0 {
		userId = ctx.Value(global.LoginUser).(models.LoginUser).ID
	}
	// 获取用户信息
	userDao, err := u.Dao.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if userDao.Avatar != "" {
		err = u.UpdateAvatar(ctx, userDao)
		if err != nil {
			return nil, err
		}
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)
	user.IsFocus = false
	// 如果登录用户不为空，且id不为0，则判断是否关注
	if ctx.Value(global.LoginUser) != nil && *idDTO.ID != 0 {
		isFocus, err := u.RelationDao.IsFocused(ctx, *idDTO.ID)
		if err != nil {
			return nil, err
		}
		user.IsFocus = isFocus
	}
	return user, nil
}

// GetUserList 获取用户列表
func (u *UserService) GetUserList(
	ctx context.Context, userListDTO *dto.UserListDTO,
) ([]*dto.User, int64, error) {
	var usersDao []*models.User
	var total int64
	var err error
	if userListDTO.Key == "" {
		usersDao, total, err = u.Dao.GetUserList(ctx, userListDTO)
		if err != nil {
			return nil, 0, err
		}
	} else {
		usersDao, total, err = u.Dao.GetUserListBySearch(ctx, userListDTO)
		if err != nil {
			return nil, 0, err
		}
	}

	if len(usersDao) == 0 {
		return nil, 0, nil
	}
	for _, user := range usersDao {
		if user.Avatar == "" {
			continue
		}
		err = u.UpdateAvatar(ctx, user)
		if err != nil {
			return nil, 0, err
		}
	}
	// 如果登录用户不为空，且id不为0，则判断是否关注
	focusMap := make(map[uint]bool)
	if ctx.Value(global.LoginUser) != nil {
		userIds := make([]uint, 0, len(usersDao))
		for _, user := range usersDao {
			userIds = append(userIds, user.ID)
		}
		focusMap, err = u.RelationDao.IsFocusedByList(ctx, userIds)
		if err != nil {
			return nil, 0, err
		}
	}
	var users = make([]*dto.User, 0, len(usersDao))
	for _, user := range usersDao {
		var userDTO = new(dto.User)
		_ = copier.Copy(userDTO, user)
		if ctx.Value(global.LoginUser) != nil && user.ID != ctx.Value(global.LoginUser).(models.LoginUser).ID {
			userDTO.IsFocus = focusMap[user.ID]
		} else {
			userDTO.IsFocus = false
		}
		users = append(users, userDTO)
	}
	return users, total, nil
}

// UpdateAvatar 更新头像
func (u *UserService) UpdateAvatar(ctx context.Context, user *models.User) error {
	ok, err := u.Dao.CheckAvatarUrl(user.Avatar)
	if err != nil {
		return err
	}
	if ok {
		avatarUrl, err := u.Dao.GetRemoteAvatarUrl(ctx, user.ID)
		if err != nil {
			return err
		}
		user.Avatar = avatarUrl

		go func(us *models.User) {
			err = u.Dao.UpdateDBUrl(ctx, us.ID, us.Avatar)
			if err != nil {
				u.logger.Error(err)
			}
		}(user)
	}
	return nil
}

// UpdateUser 更新用户信息
func (u *UserService) UpdateUser(ctx context.Context, updateDTO *dto.UserUpdateDTO) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	userDao, err := u.Dao.GetUserById(ctx, userId)
	if err != nil {
		return err
	}
	if updateDTO.Password != "" {
		ok := utils.ComparePassword(userDao.Password, updateDTO.Password)
		if !ok {
			return errors.New("old password error")
		}
		hashString, err := utils.Encrypt(updateDTO.NewPassword)
		if err != nil {
			return err
		}
		updateDTO.NewPassword = hashString
	}
	updateDTO.ToModel(userDao)
	return u.Dao.UpdateUser(ctx, userDao)
}

// UploadAvatar 上传头像
func (u *UserService) UploadAvatar(ctx context.Context, avatarDTO *dto.UploadAvatarDTO) (string, error) {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	err := u.Dao.UploadAvatar(ctx, avatarDTO)
	if err != nil {
		return "", err
	}
	avatarUrl, err := u.Dao.GetRemoteAvatarUrl(ctx, userId)
	if err != nil {
		return "", err
	}
	return avatarUrl, nil
}

// DeleteRemoteAvatar 删除远程头像
func (u *UserService) DeleteRemoteAvatar(ctx context.Context, abolish *dto.AbolishAvatarDTO) error {
	return u.Dao.DeleteRemoteAvatar(ctx, abolish.UserId)
}

// DeleteUserById 删除用户
func (u *UserService) DeleteUserById(ctx context.Context, idDTO *dto.CommonIDDTO) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	if *idDTO.ID != userId {
		return errors.New("don't have permission")
	}
	return u.Dao.DeleteUserById(ctx, userId)
}

// GetUserIncrease 获取月总日用户增长列表
func (u *UserService) GetUserIncrease(
	ctx context.Context, timeDTO *dto.UserIncreaseListDTO,
) ([]*dto.UserIncrease, error) {
	list, err := u.Dao.GetUserIncreaseList(ctx, timeDTO.Year, timeDTO.Month)
	if err != nil {
		return nil, err
	}
	var userIncreaseList = make([]*dto.UserIncrease, 0, len(list))
	_ = copier.Copy(&userIncreaseList, &list)
	return userIncreaseList, nil
}
