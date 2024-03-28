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

// IsFocused 是否已经关注
func (r *RelationDao) IsFocused(ctx context.Context, id uint) (bool, error) {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	err := r.DB.Model(&models.Relation{}).WithContext(ctx).
		Where("user_id = ? and focus_id = ?", userId, id).
		First(&models.Relation{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsFocusedByList 是否已经关注
func (r *RelationDao) IsFocusedByList(ctx context.Context, ids []uint) (map[uint]bool, error) {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	var relations []*models.Relation
	err := r.DB.Model(&models.Relation{}).WithContext(ctx).
		Where("user_id = ? and focus_id in ?", userId, ids).
		Find(&relations).Error
	if err != nil {
		return nil, err
	}
	if len(relations) == 0 {
		return nil, nil
	}
	relationMap := make(map[uint]bool)
	for _, relation := range relations {
		relationMap[relation.FocusID] = true
	}
	return relationMap, nil
}

// FocusTransaction 关注事务
func (r *RelationDao) FocusTransaction(ctx context.Context, userId uint) error {
	relation := models.Relation{
		UserID:  ctx.Value(global.LoginUser).(models.LoginUser).ID,
		FocusID: userId,
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

// UnFocusTransaction 取消关注事务
func (r *RelationDao) UnFocusTransaction(ctx context.Context, uId uint) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := r.DB.WithContext(ctx).Unscoped().
			Where("user_id = ? and focus_id = ?", ctx.Value(global.LoginUser).(models.LoginUser).ID, uId).
			Delete(&models.Relation{})
		if result.Error != nil {
			return result.Error
		}
		// 更新关注数
		if err := r.UserDao.UpdateFocusCount(ctx, userId, -1); err != nil {
			return err
		}
		// 更新粉丝数
		if err := r.UserDao.UpdateFansCount(ctx, uId, -1); err != nil {
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

// GetFocusListByUserId 根据用户id获取关注列表
func (r *RelationDao) GetFocusListByUserId(ctx context.Context, id uint) ([]*models.Relation, error) {
	var relations []*models.Relation
	err := r.DB.Model(&models.Relation{}).WithContext(ctx).Where("user_id = ?", id).Find(&relations).Error
	return relations, err
}

// GetFansListByUserId 根据用户id获取粉丝列表
func (r *RelationDao) GetFansListByUserId(ctx context.Context, id uint) ([]*models.Relation, error) {
	var relations []*models.Relation
	err := r.DB.Model(&models.Relation{}).WithContext(ctx).Where("focus_id = ?", id).Find(&relations).Error
	return relations, err
}

// GetFriendListByUserId 根据用户id获取朋友列表
func (r *RelationDao) GetFriendListByUserId(ctx context.Context, userId uint) ([]*models.Relation, error) {
	// 先获取关注列表
	relations, err := r.GetFocusListByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	if len(relations) == 0 {
		return nil, nil
	}
	var focusIds []uint
	for _, relation := range relations {
		focusIds = append(focusIds, relation.FocusID)
	}
	err = r.DB.Model(&models.Relation{}).WithContext(ctx).
		Where("focus_id = ? and user_id in ?", userId, focusIds).Find(&relations).Error
	return relations, err
}
