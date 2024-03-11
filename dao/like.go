package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var LikeDaoIns *LikeDao

type LikeDao struct {
	*BaseDao
	VideoDao *VideoDao
	UserDao  *UserDao
}

func NewLikeDao() *LikeDao {
	if LikeDaoIns == nil {
		LikeDaoIns = &LikeDao{
			BaseDao:  NewBaseDao(),
			VideoDao: NewVideoDao(),
			UserDao:  NewUserDao(),
		}
	}
	return LikeDaoIns
}

// GetLikeListByUserId 根据用户id获取点赞列表
func (l *LikeDao) GetLikeListByUserId(ctx context.Context, dto *dto.LikeListDTO) ([]*models.Like, error) {
	var likes []*models.Like
	err := l.DB.Model(&models.Like{}).WithContext(ctx).Where("user_id = ?", dto.UserID).Find(&likes).Error
	return likes, err
}

// IsLiked 是否已经点赞
func (l *LikeDao) IsLiked(ctx context.Context, videoId uint) (bool, error) {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	err := l.DB.Model(&models.Like{}).WithContext(ctx).Where("user_id = ? AND video_id = ?", userId, videoId).
		First(&models.Like{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsLikedByList 是否已经点赞
func (l *LikeDao) IsLikedByList(ctx context.Context, videoIds []uint) (map[uint]bool, error) {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	var likesDao []*models.Like
	likeMap := make(map[uint]bool)
	err := l.DB.Model(&models.Like{}).WithContext(ctx).
		Where("user_id = ? AND video_id IN ?", userId, videoIds).
		Find(&likesDao).Error
	if err != nil {
		return nil, err
	}
	if len(likesDao) == 0 {
		return nil, nil
	}
	for _, like := range likesDao {
		likeMap[like.VideoID] = true
	}
	return likeMap, nil
}

// AddLike 点赞操作
func (l *LikeDao) AddLike(ctx context.Context, dto *dto.LikeActionDTO) error {
	like := models.Like{
		UserID:  ctx.Value(global.LoginUser).(models.LoginUser).ID,
		VideoID: dto.VideoID,
	}
	isLiked, err := l.IsLiked(ctx, dto.VideoID)
	if err != nil {
		return err
	}
	if isLiked {
		return errors.New("already liked")
	}
	AuthorId, err := l.VideoDao.GetAuthorIdByVideoId(ctx, dto.VideoID)
	if err != nil {
		return err
	}
	return l.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&like).Error; err != nil {
			return err
		}
		// 更新被点赞数
		if err := l.UserDao.UpdateLikedCount(ctx, AuthorId, 1); err != nil {
			return err
		}
		// 更新点赞数
		if err := l.UserDao.UpdateLikeCount(ctx, like.UserID, 1); err != nil {
			return err
		}
		// 更新点赞数
		if err := l.VideoDao.UpdateVideoLikeCount(ctx, like.VideoID, 1); err != nil {
			return err
		}
		return nil
	})
}

// CancelLike 取消点赞操作
func (l *LikeDao) CancelLike(ctx context.Context, dto *dto.LikeActionDTO) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	isLiked, err := l.IsLiked(ctx, dto.VideoID)
	if err != nil {
		return err
	}
	if !isLiked {
		return errors.New("do not have this like relation")
	}
	AuthorId, err := l.VideoDao.GetAuthorIdByVideoId(ctx, dto.VideoID)
	if err != nil {
		return err
	}
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
		// 更新被点赞数
		if err := l.UserDao.UpdateLikedCount(ctx, AuthorId, -1); err != nil {
			return err
		}
		//更新点赞数
		if err := l.UserDao.UpdateLikeCount(ctx, userId, -1); err != nil {
			return err
		}
		// 更新视频点赞数
		if err := l.VideoDao.UpdateVideoLikeCount(ctx, dto.VideoID, -1); err != nil {
			return err
		}
		return nil
	})
}
