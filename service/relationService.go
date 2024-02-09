package service

import (
	"context"
	"errors"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var RelationServiceIns *RelationService

type RelationService struct {
	BaseService
	Dao *dao.RelationDao
}

func NewRelationService() *RelationService {
	if RelationServiceIns == nil {
		RelationServiceIns = &RelationService{
			Dao: dao.NewRelationDao(),
		}
	}
	return RelationServiceIns
}

// RelationAction 关注/取消关注
func (r *RelationService) RelationAction(ctx context.Context, dto dto.RelationActionDto) error {
	if dto.UserID == ctx.Value(global.LoginUser).(models.LoginUser).ID {
		return errors.New("can't Focus/unFocus yourself")
	}
	switch dto.ActionType {
	case 1:
		return r.Dao.Focus(ctx, dto)
	case 2:
		return r.Dao.UnFocus(ctx, dto)
	default:
		return errors.New("action type error")
	}
}

// GetFocusList 关注列表
func (r *RelationService) GetFocusList(ctx context.Context, dto dto.FocusListDto) (int64, []*models.Relation, error) {
	return r.Dao.GetFocusList(ctx, dto)
}

// GetFansList 粉丝列表
func (r *RelationService) GetFansList(ctx context.Context, dto dto.FansListDto) (int64, []*models.Relation, error) {
	return r.Dao.GetFansList(ctx, dto)
}
