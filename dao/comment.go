package dao

import (
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var CommentDaoIns *CommentDao

type CommentDao struct {
	*BaseDao
}

func NewCommentDao() *CommentDao {
	if CommentDaoIns == nil {
		CommentDaoIns = &CommentDao{
			BaseDao: NewBaseDao(),
		}
	}
	return CommentDaoIns
}

// GetCommentList 获取评论列表
func (c *CommentDao) GetCommentList(cListDTO *dto.CommentListDTO) ([]*models.Comment, int64, error) {
	var comments []*models.Comment
	var total int64
	err := c.DB.Model(&models.Comment{}).
		Scopes(Paginate(cListDTO.CommonPageDTO)).Find(&comments).
		Offset(-1).Limit(-1).Count(&total).Error
	return comments, total, err
}

// AddComment 添加评论
func (c *CommentDao) AddComment(comment *models.Comment) (*models.Comment, error) {
	err := c.DB.Create(comment).Error
	return comment, err
}

// GetCommentById 根据id获取评论
func (c *CommentDao) GetCommentById(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := c.DB.Model(&models.Comment{}).First(&comment, id).Error
	return &comment, err
}

// DeleteCommentById 根据id删除评论
func (c *CommentDao) DeleteCommentById(id uint) error {
	return c.DB.Delete(&models.Comment{}, id).Error
}
