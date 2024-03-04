package dao

import (
	"context"
	"errors"

	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
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
func (l *CollectDao) GetCollectListByUserId(ctx context.Context, dto *dto.CollectListDTO) ([]*models.Collect, error) {
	var collects []*models.Collect
	err := l.DB.Model(&models.Collect{}).WithContext(ctx).Where("user_id = ?", dto.UserID).Find(&collects).Error
	return collects, err
}

// isCollected 是否已经收藏
func (l *CollectDao) isCollected(ctx context.Context, collect models.Collect) bool {
	err := l.DB.Model(&models.Collect{}).WithContext(ctx).
		First(&collect).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// AddCollect 收藏操作
func (l *CollectDao) AddCollect(ctx context.Context, dto *dto.CollectActionDTO) error {
	collect := models.Collect{
		UserID:  ctx.Value(global.LoginUser).(models.LoginUser).ID,
		VideoID: dto.VideoID,
	}
	if l.isCollected(ctx, collect) {
		return errors.New("already collected")
	}
	return l.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&collect).Error; err != nil {
			return err
		}
		// 更新点赞数
		if err := l.UserDao.UpdateCollectCount(ctx, collect.UserID, 1); err != nil {
			return err
		}
		// 更新点赞数
		if err := l.VideoDao.UpdateVideoCollectCount(ctx, collect.VideoID, 1); err != nil {
			return err
		}
		return nil
	})
}

// CancelCollect 取消收藏操作
func (l *CollectDao) CancelCollect(ctx context.Context, dto *dto.CollectActionDTO) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	return l.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.WithContext(ctx).Unscoped().
			Where("user_id = ? AND video_id = ?", userId, dto.VideoID).
			Delete(&models.Like{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("do not have this like relation")
		}
		//更新点赞数
		if err := l.UserDao.UpdateCollectCount(ctx, userId, -1); err != nil {
			return err
		}
		// 更新视频点赞数
		if err := l.VideoDao.UpdateVideoCollectCount(ctx, dto.VideoID, -1); err != nil {
			return err
		}
		return nil
	})
}
