package dao

import (
	"context"
	"errors"

	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"gorm.io/gorm"
)

var CollectDaoIns *CollectDao

type CollectDao struct {
	*BaseDao
	VideoDao   *VideoDao
	UserDao    *UserDao
	CollectDao *CollectDao
}

func NewCollectDao() *CollectDao {
	if CollectDaoIns == nil {
		CollectDaoIns = &CollectDao{
			BaseDao:  NewBaseDao(),
			VideoDao: NewVideoDao(),
			UserDao:  NewUserDao(),
		}
	}
	return CollectDaoIns
}

// GetCollectListByUserId 根据用户id获取收藏列表
func (l *CollectDao) GetCollectListByUserId(ctx context.Context, userId uint) ([]*models.Collect, error) {
	var collects []*models.Collect
	err := l.DB.Model(&models.Collect{}).WithContext(ctx).Where("user_id = ?", userId).Find(&collects).Error
	return collects, err
}

// IsCollected 是否已经收藏
func (l *CollectDao) IsCollected(ctx context.Context, videoId uint) (bool, error) {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	err := l.DB.Model(&models.Collect{}).WithContext(ctx).Where("user_id = ? AND video_id = ?", userId, videoId).
		First(&models.Collect{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsCollectedByList 是否已经收藏
func (l *CollectDao) IsCollectedByList(ctx context.Context, videoIds []uint) (map[uint]bool, error) {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	var collectsDao []*models.Collect
	collectMap := make(map[uint]bool)
	err := l.DB.Model(&models.Collect{}).WithContext(ctx).
		Where("user_id = ? AND video_id IN ?", userId, videoIds).
		Find(&collectsDao).Error
	if err != nil {
		return nil, err
	}
	if len(collectsDao) == 0 {
		return nil, nil
	}
	for _, collect := range collectsDao {
		collectMap[collect.VideoID] = true
	}
	return collectMap, nil
}

// CreateCollectTransaction 创建收藏事务
func (l *CollectDao) CreateCollectTransaction(ctx context.Context, userId, videoId uint) error {
	collect := models.Collect{
		UserID:  userId,
		VideoID: videoId,
	}
	return l.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&collect).Error; err != nil {
			return err
		}
		// 更新收藏数
		if err := l.UserDao.UpdateCollectCount(ctx, collect.UserID, 1); err != nil {
			return err
		}
		// 更新收藏数
		if err := l.VideoDao.UpdateVideoCollectCount(ctx, collect.VideoID, 1); err != nil {
			return err
		}
		return nil
	})
}

// DeleteCollectTransaction 删除收藏事务
func (l *CollectDao) DeleteCollectTransaction(ctx context.Context, userId, videoId uint) error {
	return l.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.WithContext(ctx).Unscoped().
			Where("user_id = ? AND video_id = ?", userId, videoId).
			Delete(&models.Collect{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("do not have this collect relation")
		}
		//更新收藏数
		if err := l.UserDao.UpdateCollectCount(ctx, userId, -1); err != nil {
			return err
		}
		// 更新视频收藏数
		if err := l.VideoDao.UpdateVideoCollectCount(ctx, videoId, -1); err != nil {
			return err
		}
		return nil
	})
}
