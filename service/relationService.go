package service

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var RelationServiceIns *RelationService

type RelationService struct {
	BaseService
	UserDao *dao.UserDao
	Dao     *dao.RelationDao
}

func NewRelationService() *RelationService {
	if RelationServiceIns == nil {
		RelationServiceIns = &RelationService{
			Dao:     dao.NewRelationDao(),
			UserDao: dao.NewUserDao(),
		}
	}
	return RelationServiceIns
}

// RelationAction 关注/取消关注
func (r *RelationService) RelationAction(ctx context.Context, relationDto dto.RelationActionDto) error {
	if !r.UserDao.IsExist(ctx, relationDto.UserID) {
		return errors.New("user not exist")
	}

	if relationDto.UserID == ctx.Value(global.LoginUser).(models.LoginUser).ID {
		return errors.New("can't Focus/unFocus yourself")
	}
	switch relationDto.ActionType {
	case 1:
		return r.Dao.Focus(ctx, relationDto)
	case 2:
		return r.Dao.UnFocus(ctx, relationDto)
	default:
		return errors.New("action type error")
	}
}

// GetFocusList 关注列表
func (r *RelationService) GetFocusList(ctx context.Context, fListDto dto.FocusListDto) (int64, []*dto.User, error) {

	total, focusList, err := r.Dao.GetFocusList(ctx, fListDto)
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return 0, nil, nil
	}
	var focusUserIDs []uint
	for _, focus := range focusList {
		focusUserIDs = append(focusUserIDs, focus.FocusID)
	}
	userList, err := r.UserDao.GetUserListByIds(ctx, focusUserIDs)
	if err != nil {
		return 0, nil, err
	}
	var users []*dto.User
	_ = copier.Copy(&users, &userList)
	return total, users, nil
}

// GetFansList 粉丝列表
func (r *RelationService) GetFansList(ctx context.Context, fansListDto dto.FansListDto) (int64, []*dto.User, error) {
	total, fansList, err := r.Dao.GetFansList(ctx, fansListDto)
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return 0, nil, nil
	}
	var fansUserIDs []uint
	for _, fans := range fansList {
		fansUserIDs = append(fansUserIDs, fans.UserID)
	}
	userList, err := r.UserDao.GetUserListByIds(ctx, fansUserIDs)
	if err != nil {
		return 0, nil, err
	}
	var users []*dto.User
	_ = copier.Copy(&users, &userList)
	return total, users, nil
}
