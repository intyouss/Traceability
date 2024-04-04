package dao

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/intyouss/Traceability/utils"

	"github.com/intyouss/Traceability/global"
	"github.com/minio/minio-go/v7"

	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
	"gorm.io/gorm"
)

var UserDaoIns *UserDao

type UserDao struct {
	*BaseDao
	OSS *utils.MinioClient
}

func NewUserDao() *UserDao {
	if UserDaoIns == nil {
		UserDaoIns = &UserDao{
			BaseDao: NewBaseDao(),
			OSS:     global.OSS,
		}
	}
	return UserDaoIns
}

func (u *UserDao) GetUserByName(ctx context.Context, username string) (*models.User, error) {
	var user *models.User
	err := u.DB.Model(&models.User{}).WithContext(ctx).Where("username = ?", username).First(&user).Error
	return user, err
}

// GetUserByNameAndPassword 根据用户名和密码获取用户
func (u *UserDao) GetUserByNameAndPassword(username, password string) (*models.User, error) {
	var user *models.User
	err := u.DB.Model(&models.User{}).
		Where("username = ? AND password = ?", username, password).
		First(&user).Error
	return user, err
}

// CheckUserNameExist 检查用户名是否存在
func (u *UserDao) CheckUserNameExist(username string) bool {
	var user models.User
	err := u.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// AddUser 添加用户
func (u *UserDao) AddUser(ctx context.Context, userAddDTO *dto.UserAddDTO) (*models.User, error) {
	user := &models.User{}
	userAddDTO.ToModel(user)
	err := u.DB.WithContext(ctx).Create(&user).Error
	if err == nil {
		userAddDTO.ID = &user.ID
		userAddDTO.Password = ""
	}
	return user, err
}

// GetUserById 根据id获取用户
func (u *UserDao) GetUserById(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&models.User{}).WithContext(ctx).First(&user, id).Error
	return &user, err
}

// GetUserIdsBySearchKey 根据关键字模糊搜索用户id列表
func (u *UserDao) GetUserIdsBySearchKey(ctx context.Context, key string) ([]uint, error) {
	var ids []uint
	err := u.DB.Model(&models.User{}).WithContext(ctx).
		Where("username like ?", "%"+key+"%").
		Pluck("id", &ids).Error
	return ids, err
}

// GetUserListBySearch 模糊搜索用户列表
func (u *UserDao) GetUserListBySearch(
	ctx context.Context, userListDto *dto.UserListDTO,
) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64
	err := u.DB.Model(&models.User{}).WithContext(ctx).
		Where("username like ?", "%"+userListDto.Key+"%").
		Scopes(Paginate(userListDto.CommonPageDTO)).Find(&users).
		Offset(-1).Limit(-1).Count(&total).Error
	return users, total, err
}

// GetUserList 获取用户列表
func (u *UserDao) GetUserList(
	ctx context.Context, userListDto *dto.UserListDTO,
) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64
	err := u.DB.Model(&models.User{}).WithContext(ctx).
		Scopes(Paginate(userListDto.CommonPageDTO)).Find(&users).
		Offset(-1).Limit(-1).Count(&total).Error
	return users, total, err
}

// UpdateUser 更新用户信息
func (u *UserDao) UpdateUser(ctx context.Context, user *models.User) error {
	return u.DB.WithContext(ctx).Updates(&user).Error
}

// DeleteUserById 根据id删除用户
func (u *UserDao) DeleteUserById(ctx context.Context, id uint) error {
	return u.DB.WithContext(ctx).Delete(&models.User{}, id).Error
}

// GetUserListByIds 根据id列表获取用户列表
func (u *UserDao) GetUserListByIds(ctx context.Context, ids []uint) ([]*models.User, error) {
	var users []*models.User
	err := u.DB.Model(&models.User{}).WithContext(ctx).Where("id in (?)", ids).Find(&users).Error
	return users, err
}

// IsExist 用户是否存在
func (u *UserDao) IsExist(ctx context.Context, id uint) bool {
	err := u.DB.Model(&models.User{}).WithContext(ctx).First(&models.User{}, id).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// UploadAvatar 上传头像
func (u *UserDao) UploadAvatar(ctx context.Context, upload *dto.UploadAvatarDTO) error {
	// 读取头像
	avatarSize := upload.AvatarData.Size
	avatarData, err := upload.AvatarData.Open()
	if err != nil {
		return err
	}
	defer avatarData.Close()
	// 上传头像
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	fileName := "avatars/" + strconv.Itoa(int(userId)) + ".png"
	err = u.OSS.UploadSizeFile(
		ctx, VideoBucket, fileName, avatarData, avatarSize, minio.PutObjectOptions{
			ContentType: "image/png",
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRemoteAvatar 删除远程头像
func (u *UserDao) DeleteRemoteAvatar(ctx context.Context, userId uint) error {
	avatarFileName := "avatars/" + strconv.Itoa(int(userId)) + ".png"
	return u.OSS.RemoveFile(ctx, VideoBucket, avatarFileName)
}

// GetRemoteAvatarUrl 获取远程头像url
func (u *UserDao) GetRemoteAvatarUrl(ctx context.Context, userId uint) (avatarURL string, err error) {
	hours, days := 24, 7
	fileName := "avatars/" + strconv.Itoa(int(userId)) + ".png"
	timeDura := time.Hour * time.Duration(hours*days)
	urls, err := u.OSS.GetFileURL(ctx, "oss", fileName, timeDura)
	if err != nil {
		return "", err
	}
	avatarURL = urls.String()
	return
}

// CheckAvatarUrl 检查用户头像url是否失效
func (u *UserDao) CheckAvatarUrl(avatarUrl string) (bool, error) {
	ok, err := u.OSS.CheckUrl(avatarUrl)
	if err != nil {
		return false, err
	}
	return ok, nil
}

// UpdateDBUrl 更新数据库url
func (u *UserDao) UpdateDBUrl(ctx context.Context, userId uint, avatarUrl string) error {
	return u.DB.WithContext(ctx).Where("id = ?", userId).
		Updates(&models.User{Avatar: avatarUrl}).Error
}

// GetUserIncrease 获取用户日增长记录
func (u *UserDao) GetUserIncrease(ctx context.Context, year, month, day uint) (bool, *models.UserIncrease, error) {
	var userIncrease models.UserIncrease
	err := u.DB.Model(&models.UserIncrease{}).WithContext(ctx).
		Where("year = ? and month = ? and day = ?", year, month, day).
		FirstOrCreate(&userIncrease).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil, nil
	}
	if err != nil {
		return false, nil, err
	}
	return true, &userIncrease, err
}

// GetUserIncreaseList 获取用户增长记录列表
func (u *UserDao) GetUserIncreaseList(ctx context.Context, year, month uint) ([]*models.UserIncrease, error) {
	var userIncreases []*models.UserIncrease
	err := u.DB.Model(&models.UserIncrease{}).WithContext(ctx).Where("year = ? and month = ?", year, month).
		Find(&userIncreases).Error
	return userIncreases, err
}

// UpdateUserIncreaseCount 更新用户日增长记录
func (u *UserDao) UpdateUserIncreaseCount(ctx context.Context, year, month, day uint, count int) error {
	value := map[string]interface{}{"count": gorm.Expr("count + ?", count)}
	return u.DB.Model(&models.UserIncrease{}).WithContext(ctx).
		Where("month = ? and day = ? and year = ?", month, day, year).
		Updates(value).Error
}

// UpdateFocusCount 更新关注数
func (u *UserDao) UpdateFocusCount(ctx context.Context, id uint, count int) error {
	value := map[string]interface{}{"focus_count": gorm.Expr("focus_count + ?", count)}
	return u.DB.Model(&models.User{}).WithContext(ctx).Where("id = ?", id).
		Updates(value).Error
}

// UpdateFansCount 更新粉丝数
func (u *UserDao) UpdateFansCount(ctx context.Context, id uint, count int) error {
	value := map[string]interface{}{"fans_count": gorm.Expr("fans_count + ?", count)}
	return u.DB.Model(&models.User{}).WithContext(ctx).Where("id = ?", id).
		Updates(value).Error
}

// UpdateLikeCount 更新点赞数
func (u *UserDao) UpdateLikeCount(ctx context.Context, id uint, count int) error {
	value := map[string]interface{}{"like_count": gorm.Expr("like_count + ?", count)}
	return u.DB.Model(&models.User{}).WithContext(ctx).Where("id = ?", id).
		Updates(value).Error
}

// UpdateLikedCount 更新被点赞数
func (u *UserDao) UpdateLikedCount(ctx context.Context, id uint, count int) error {
	value := map[string]interface{}{"liked_count": gorm.Expr("liked_count + ?", count)}
	return u.DB.Model(&models.User{}).WithContext(ctx).Where("id = ?", id).
		Updates(value).Error
}

// UpdateVideoCount 更新视频数
func (u *UserDao) UpdateVideoCount(ctx context.Context, id uint, count int) error {
	value := map[string]interface{}{"video_count": gorm.Expr("video_count + ?", count)}
	return u.DB.Model(&models.User{}).WithContext(ctx).Where("id = ?", id).
		Updates(value).Error
}

// UpdateCollectCount 更新收藏数
func (u *UserDao) UpdateCollectCount(ctx context.Context, id uint, count int) error {
	value := map[string]interface{}{"collect_count": gorm.Expr("collect_count + ?", count)}
	return u.DB.Model(&models.User{}).WithContext(ctx).Where("id = ?", id).
		Updates(value).Error
}
