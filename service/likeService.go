package service

import (
	"context"
	"errors"
	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var LikeServiceIns *LikeService

type LikeService struct {
	BaseService
	Dao *dao.LikeDao
}

func NewLikeService() *LikeService {
	if LikeServiceIns == nil {
		LikeServiceIns = &LikeService{
			Dao: dao.NewLikeDao(),
		}
	}
	return LikeServiceIns
}

// GetLikeList 获取点赞列表
func (l *LikeService) GetLikeList(ctx context.Context, dto *dto.LikeListDTO) ([]*models.Like, error) {
	myUserID := ctx.Value(global.LoginUser).(uint)
	if myUserID != dto.UserID {
		return nil, errors.New("don't have permission")
	}
	return l.Dao.GetLikeListByUserId(ctx, dto)
}

// LikeAction 喜爱操作
func (l *LikeService) LikeAction(ctx context.Context, dto *dto.LikeActionDTO) error {
	switch dto.ActionType {
	case 1:
		return l.Dao.AddLike(ctx, dto)
	case 2:
		return l.Dao.CancelLike(ctx, dto)
	default:
		return errors.New("action type error")
	}
}
