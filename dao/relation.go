package dao

import (
	"context"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var RelationDaoIns *RelationDao

type RelationDao struct {
	*BaseDao
}

func NewRelationDao() *RelationDao {
	if RelationDaoIns == nil {
		RelationDaoIns = &RelationDao{
			BaseDao: NewBaseDao(),
		}
	}
	return RelationDaoIns
}

// Focus 关注
func (r *RelationDao) Focus(ctx context.Context, dto dto.RelationActionDto) error {
	relation := &models.Relation{
		UserID:  ctx.Value(global.LoginUser).(uint),
		FocusID: dto.UserId,
	}
	return r.DB.Model(&models.Relation{}).Create(&relation).Error
}

// UnFocus 取消关注
func (r *RelationDao) UnFocus(ctx context.Context, dto dto.RelationActionDto) error {
	return r.DB.Model(&models.Relation{}).WithContext(ctx).
		Where("user_id = ? and focus_id = ?", ctx.Value(global.LoginUser), dto.UserId).
		Delete(&models.Relation{}).Error
}

// GetFocusList 关注列表
func (r *RelationDao) GetFocusList(ctx context.Context, dto dto.FocusListDto) ([]*models.Relation, error) {
	var relations []*models.Relation
	err := r.DB.Model(&models.Relation{}).WithContext(ctx).
		Where("user_id = ?", dto.UserId).Find(&relations).Error
	return relations, err
}

// GetFansList 粉丝列表
func (r *RelationDao) GetFansList(ctx context.Context, dto dto.FansListDto) ([]*models.Relation, error) {
	var relations []*models.Relation
	err := r.DB.Model(&models.Relation{}).WithContext(ctx).
		Where("focus_id = ?", dto.UserId).Find(&relations).Error
	return relations, err
}
