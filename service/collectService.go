package service

import (
	"context"
	"errors"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var CollectServiceIns *CollectService

type CollectService struct {
	BaseService
	Dao *dao.CollectDao
}

func NewCollectService() *CollectService {
	if CollectServiceIns == nil {
		CollectServiceIns = &CollectService{
			Dao: dao.NewCollectDao(),
		}
	}
	return CollectServiceIns
}

// GetCollectList 获取收藏列表
func (l *CollectService) GetCollectList(ctx context.Context, dto *dto.CollectListDTO) ([]*models.Collect, error) {
	myUserID := ctx.Value(global.LoginUser).(models.LoginUser).ID
	if myUserID != dto.UserID {
		return nil, errors.New("don't have permission")
	}
	return l.Dao.GetCollectListByUserId(ctx, dto)
}

// CollectAction 收藏操作
func (l *CollectService) CollectAction(ctx context.Context, dto *dto.CollectActionDTO) error {
	switch dto.ActionType {
	case 1:
		return l.Dao.AddCollect(ctx, dto)
	case 2:
		return l.Dao.CancelCollect(ctx, dto)
	default:
		return errors.New("action type error")
	}
}
