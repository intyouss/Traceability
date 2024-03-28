package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var CommentDaoIns *CommentDao

type CommentDao struct {
	*BaseDao
	VideoDao *VideoDao
}

func NewCommentDao() *CommentDao {
	if CommentDaoIns == nil {
		CommentDaoIns = &CommentDao{
			BaseDao:  NewBaseDao(),
			VideoDao: NewVideoDao(),
		}
	}
	return CommentDaoIns
}

// GetCommentList 获取评论列表
func (c *CommentDao) GetCommentList(
	ctx context.Context, pageDto dto.CommonPageDTO, videoId uint,
) ([]*models.Comment, int64, error) {
	var comments []*models.Comment
	var total int64
	err := c.DB.Model(&models.Comment{}).WithContext(ctx).
		Scopes(Paginate(pageDto)).Where("video_id = ?", videoId).Find(&comments).
		Offset(-1).Limit(-1).Count(&total).Error
	return comments, total, err
}

// AddComment 添加评论
func (c *CommentDao) AddComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	err := c.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(comment).Error; err != nil {
			return err
		}
		// 更新评论数
		if err := c.VideoDao.UpdateCommentCount(ctx, comment.VideoId, 1); err != nil {
			return err
		}
		return nil
	})
	return comment, err
}

// GetCommentById 根据id获取评论
func (c *CommentDao) GetCommentById(ctx context.Context, id uint) (*models.Comment, error) {
	var comment models.Comment
	err := c.DB.Model(&models.Comment{}).WithContext(ctx).First(&comment, id).Error
	return &comment, err
}

// DeleteCommentById 根据id删除评论
func (c *CommentDao) DeleteCommentById(ctx context.Context, id uint) error {
	return c.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var comment models.Comment
		if err := tx.WithContext(ctx).First(&comment, id).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Unscoped().Delete(&models.Comment{}, id).Error; err != nil {
			return err
		}
		// 更新评论数
		if err := c.VideoDao.UpdateCommentCount(ctx, comment.VideoId, -1); err != nil {
			return err
		}
		return nil
	})
}
