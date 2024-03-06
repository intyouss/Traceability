package service

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"

	"github.com/intyouss/Traceability/global"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var CommentServiceIns *CommentService

type CommentService struct {
	BaseService
	VideoDao *dao.VideoDao
	UserDao  *dao.UserDao
	Dao      *dao.CommentDao
}

func NewCommentService() *CommentService {
	if CommentServiceIns == nil {
		CommentServiceIns = &CommentService{
			Dao:      dao.NewCommentDao(),
			VideoDao: dao.NewVideoDao(),
			UserDao:  dao.NewUserDao(),
		}
	}
	return CommentServiceIns
}

// GetCommentList 获取评论列表
func (c *CommentService) GetCommentList(
	ctx context.Context, cListDTO *dto.CommentListDTO,
) ([]*dto.Comment, int64, error) {
	if !c.VideoDao.IsExist(ctx, cListDTO.VideoID) {
		return nil, 0, errors.New("video not exist")
	}

	commentsDao, total, err := c.Dao.GetCommentList(ctx, cListDTO)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}

	commentUserMap := make(map[uint]*models.User)
	for _, comment := range commentsDao {
		commentUserMap[comment.UserId] = nil
	}

	var userIds []uint
	for userId := range commentUserMap {
		userIds = append(userIds, userId)
	}

	users, err := c.UserDao.GetUserListByIds(ctx, userIds)
	if err != nil {
		return nil, 0, err
	}

	for _, user := range users {
		commentUserMap[user.ID] = user
	}

	var comments = make([]*dto.Comment, len(commentsDao))
	for i, comment := range commentsDao {
		comments[i].ID = comment.ID
		comments[i].Content = comment.Content
		comments[i].CreatedAt = comment.CreatedAt.Format("2006-01-02 15:04:05")
		var user = new(dto.User)
		_ = copier.Copy(user, commentUserMap[comment.UserId])
		comments[i].User = user
	}
	return comments, total, nil
}

// AddComment 添加评论
func (c *CommentService) AddComment(ctx context.Context, cAddDTO *dto.AddCommentDTO) (*dto.Comment, error) {
	if !c.VideoDao.IsExist(ctx, cAddDTO.VideoId) {
		return nil, errors.New("video not exist")
	}

	var commentDao models.Comment
	cAddDTO.ToModel(&commentDao)
	commentDao.UserId = ctx.Value(global.LoginUser).(models.LoginUser).ID
	cm, err := c.Dao.AddComment(ctx, &commentDao)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	userDao, err := c.UserDao.GetUserById(ctx, cm.UserId)
	if err != nil {
		return nil, err
	}

	var comment = new(dto.Comment)
	_ = copier.Copy(comment, cm)
	comment.CreatedAt = cm.CreatedAt.Format("2006-01-02 15:04:05")
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)
	comment.User = user
	return comment, nil

}

// DeleteCommentById 根据id删除评论
func (c *CommentService) DeleteCommentById(ctx context.Context, cDeleteDTO *dto.DeleteCommentDTO) error {
	video, err := c.Dao.GetCommentById(ctx, *cDeleteDTO.ID)
	if err != nil {
		return err
	}
	if video.UserId != ctx.Value(global.LoginUser).(models.LoginUser).ID {
		return errors.New("permission denied")
	}
	return c.Dao.DeleteCommentById(ctx, *cDeleteDTO.ID)
}
