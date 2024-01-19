package service

import (
	"errors"
	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var CommentServiceIns *CommentService

type CommentService struct {
	BaseService
	Dao *dao.CommentDao
}

func NewCommentService() *CommentService {
	if CommentServiceIns == nil {
		CommentServiceIns = &CommentService{
			Dao: dao.NewCommentDao(),
		}
	}
	return CommentServiceIns
}

// GetCommentList 获取评论列表
func (c *CommentService) GetCommentList(cListDTO *dto.CommentListDTO) ([]*models.Comment, int64, error) {
	return c.Dao.GetCommentList(cListDTO)
}

// AddComment 添加评论
func (c *CommentService) AddComment(cActionDTO *dto.CommentActionDTO) (*models.Comment, error) {
	var comment models.Comment
	cActionDTO.AddActionToModel(&comment)
	return c.Dao.AddComment(&comment)
}

// DeleteCommentById 根据id删除评论
func (c *CommentService) DeleteCommentById(cActionDTO *dto.CommentActionDTO) error {
	return c.Dao.DeleteCommentById(cActionDTO.CommentID)
}

// CommentAction 评论操作
func (c *CommentService) CommentAction(cActionDTO *dto.CommentActionDTO) (*models.Comment, error) {
	var comment *models.Comment
	var err error
	switch cActionDTO.ActionType {
	case 1:
		comment, err = c.AddComment(cActionDTO)
	case 2:
		err = c.DeleteCommentById(cActionDTO)
	default:
		err = errors.New("invalid action type")
	}
	return comment, err
}
