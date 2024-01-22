package service

import (
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
func (c *CommentService) AddComment(cAddDTO *dto.AddCommentDTO) (*models.Comment, error) {
	var comment models.Comment
	cAddDTO.ToModel(&comment)
	return c.Dao.AddComment(&comment)
}

// DeleteCommentById 根据id删除评论
func (c *CommentService) DeleteCommentById(cDeleteDTO *dto.DeleteCommentDTO) error {
	return c.Dao.DeleteCommentById(cDeleteDTO.CommentID)
}
