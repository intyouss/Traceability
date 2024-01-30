package service

import (
	"context"
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
func (c *CommentService) GetCommentList(
	ctx context.Context, cListDTO *dto.CommentListDTO,
) ([]*models.Comment, int64, error) {
	return c.Dao.GetCommentList(ctx, cListDTO)
}

// AddComment 添加评论
func (c *CommentService) AddComment(ctx context.Context, cAddDTO *dto.AddCommentDTO) (*models.Comment, error) {
	var comment models.Comment
	cAddDTO.ToModel(&comment)
	return c.Dao.AddComment(ctx, &comment)
}

// DeleteCommentById 根据id删除评论
func (c *CommentService) DeleteCommentById(ctx context.Context, cDeleteDTO *dto.DeleteCommentDTO) error {
	return c.Dao.DeleteCommentById(ctx, cDeleteDTO.CommentID)
}
