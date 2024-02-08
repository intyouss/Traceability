package dao

import (
	"context"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var LikeDaoIns *LikeDao

type LikeDao struct {
	*BaseDao
}

func NewLikeDao() *LikeDao {
	if LikeDaoIns == nil {
		LikeDaoIns = &LikeDao{
			BaseDao: NewBaseDao(),
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

// AddLike 点赞操作
func (l *LikeDao) AddLike(ctx context.Context, dto *dto.LikeActionDTO) error {
	like := models.Like{
		UserID:  ctx.Value(global.LoginUser).(models.LoginUser).ID,
		VideoID: dto.VideoID,
	}
	return l.DB.Model(&models.Like{}).WithContext(ctx).Create(&like).Error
}

// CancelLike 取消点赞操作
func (l *LikeDao) CancelLike(ctx context.Context, dto *dto.LikeActionDTO) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	return l.DB.Model(&models.Like{}).WithContext(ctx).
		Where("user_id = ? AND video_id = ?", userId, dto.VideoID).
		Delete(&models.Like{}).Error
}
