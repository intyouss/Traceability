package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var RelationDaoIns *RelationDao

type RelationDao struct {
	*BaseDao
	UserDao *UserDao
}

func NewRelationDao() *RelationDao {
	if RelationDaoIns == nil {
		RelationDaoIns = &RelationDao{
			BaseDao: NewBaseDao(),
			UserDao: NewUserDao(),
		}
	}
	return RelationDaoIns
}

// isFocused 是否已经关注
func (r *RelationDao) isFocused(ctx context.Context, relation models.Relation) bool {
	err := r.DB.Model(&models.Relation{}).WithContext(ctx).First(&relation).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Focus 关注
func (r *RelationDao) Focus(ctx context.Context, dto dto.RelationActionDto) error {
	relation := models.Relation{
		UserID:  ctx.Value(global.LoginUser).(models.LoginUser).ID,
		FocusID: dto.UserID,
	}
	if r.isFocused(ctx, relation) {
		return errors.New("already focused")
	}
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&relation).Error; err != nil {
			return err
		}
		// 更新关注数
		if err := r.UserDao.UpdateFocusCount(ctx, relation.UserID, 1); err != nil {
			return err
		}
		// 更新粉丝数
		if err := r.UserDao.UpdateFansCount(ctx, relation.FocusID, 1); err != nil {
			return err
		}
		return nil
	})
}

// UnFocus 取消关注
func (r *RelationDao) UnFocus(ctx context.Context, dto dto.RelationActionDto) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := r.DB.WithContext(ctx).Unscoped().
			Where("user_id = ? and focus_id = ?", ctx.Value(global.LoginUser).(models.LoginUser).ID, dto.UserID).
			Delete(&models.Relation{})
		if result.Error != nil {
			return result.Error
		}
		// 如果没有删除任何记录，说明没有这个关系
		if result.RowsAffected == 0 {
			return errors.New("do not have this focus relation")
		}
		// 更新关注数
		if err := r.UserDao.UpdateFocusCount(ctx, userId, -1); err != nil {
			return err
		}

		// 更新粉丝数
		if err := r.UserDao.UpdateFansCount(ctx, dto.UserID, -1); err != nil {
			return err
		}
		return nil
	})
}

// GetFocusList 关注列表
func (r *RelationDao) GetFocusList(ctx context.Context, dto dto.FocusListDto) (int64, []*models.Relation, error) {
	var relations []*models.Relation
	var total int64
	err := r.DB.Model(&models.Relation{}).WithContext(ctx).
		Scopes(Paginate(dto.CommonPageDTO)).Where("user_id = ?", dto.UserID).Find(&relations).
		Offset(-1).Limit(-1).Count(&total).Error
	return total, relations, err
}

// GetFansList 粉丝列表
func (r *RelationDao) GetFansList(ctx context.Context, dto dto.FansListDto) (int64, []*models.Relation, error) {
	var relations []*models.Relation
	var total int64
	err := r.DB.Model(&models.Relation{}).WithContext(ctx).
		Scopes(Paginate(dto.CommonPageDTO)).Where("focus_id = ?", dto.UserID).Find(&relations).
		Offset(-1).Limit(-1).Count(&total).Error
	return total, relations, err
}
