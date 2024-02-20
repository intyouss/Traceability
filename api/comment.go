package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
	"github.com/jinzhu/copier"
)

const (
	ErrCodeGetCommentList = iota + 20001
	ErrCodeAddComment
	ErrCodeDeleteComment
)

type CommentApi struct {
	BaseApi
	UserApi  UserApi
	VideoApi VideoApi
	Service  *service.CommentService
}

func NewCommentApi() CommentApi {
	return CommentApi{
		BaseApi:  NewBaseApi(),
		UserApi:  NewUserApi(),
		VideoApi: NewVideoApi(),
		Service:  service.NewCommentService(),
	}
}

// GetCommentList 获取评论列表
// @Summary 获取评论列表
// @Description 获取评论列表
// @Param page formData int false "页码"
// @Param limit formData int false "每页数量"
// @Param video_id formData int true "视频id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/public/comment/list [get]
func (c CommentApi) GetCommentList(ctx *gin.Context) {
	var cListDto dto.CommentListDTO
	if err := c.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &cListDto}).GetError(); err != nil {
		c.Logger.Error(err)
		c.Fail(&Response{Code: ErrCodeGetCommentList, Msg: err.Error()})
		return
	}

	if !c.VideoApi.Service.IsExist(ctx, cListDto.VideoID) {
		c.Fail(&Response{Code: ErrCodeGetCommentList, Msg: "video not exist"})
		return
	}

	commentsDao, total, err := c.Service.GetCommentList(ctx, &cListDto)
	if err != nil {
		c.Logger.Error(err)
		c.Fail(&Response{Code: ErrCodeGetCommentList, Msg: err.Error()})
		return
	}
	if total == 0 {
		c.Success(&Response{
			Data: gin.H{
				"comment": []dto.Comment{},
			},
			Total: 0,
		})
		return
	}

	commentUserMap := make(map[uint]*models.User)
	for _, comment := range commentsDao {
		commentUserMap[comment.UserId] = nil
	}

	var userIds []uint
	for userId := range commentUserMap {
		userIds = append(userIds, userId)
	}

	users, err := c.UserApi.Service.GetUserListByIds(ctx, userIds)
	if err != nil {
		c.Logger.Error(err)
		c.Fail(&Response{Code: ErrCodeGetCommentList, Msg: err.Error()})
		return
	}

	for _, user := range users {
		commentUserMap[user.ID] = user
	}

	var comments = make([]dto.Comment, len(commentsDao))
	for i, comment := range commentsDao {
		comments[i].ID = comment.ID
		comments[i].Content = comment.Content
		comments[i].CreatedAt = comment.CreatedAt.Format("2006-01-02 15:04:05")
		var user = new(dto.User)
		_ = copier.Copy(user, commentUserMap[comment.UserId])
		comments[i].User = user
	}

	c.Success(&Response{
		Data: gin.H{
			"comments": comments,
		},
		Total: total,
	})
}

// AddComment 添加评论
// @Summary 添加评论
// @Description 添加评论
// @Param token header string true "token"
// @Param video_id formData int true "视频id"
// @Param content formData string false "评论内容"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/comment/add [post]
func (c CommentApi) AddComment(ctx *gin.Context) {
	var cAddDto dto.AddCommentDTO
	if err := c.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &cAddDto}).GetError(); err != nil {
		c.Logger.Error(err)
		c.Fail(&Response{Code: ErrCodeAddComment, Msg: err.Error()})
		return
	}

	if !c.VideoApi.Service.IsExist(ctx, cAddDto.VideoID) {
		c.Fail(&Response{Code: ErrCodeGetCommentList, Msg: "video not exist"})
		return
	}

	// 添加评论
	commentDao, err := c.Service.AddComment(ctx, &cAddDto)
	if err != nil {
		c.Logger.Error(err)
		c.Fail(&Response{Code: ErrCodeAddComment, Msg: err.Error()})
		return
	}

	// 获取用户信息
	var IdDTO dto.CommonUserIDDTO
	IdDTO.ID = int(commentDao.UserId)
	userDao, err := c.UserApi.Service.GetUserById(ctx, &IdDTO)
	if err != nil {
		c.Logger.Error(err)
		c.Fail(&Response{Code: ErrCodeAddComment, Msg: err.Error()})
		return
	}

	var comment = new(dto.Comment)
	_ = copier.Copy(comment, commentDao)
	comment.CreatedAt = commentDao.CreatedAt.Format("2006-01-02 15:04:05")
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)
	comment.User = user

	c.Success(&Response{
		Data: gin.H{
			"comment": comment,
		},
	})
}

// DeleteComment 删除评论
// @Summary 删除评论
// @Description 删除评论
// @Param token header string true "token"
// @Param comment_id formData int true "评论id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/comment/delete [post]
func (c CommentApi) DeleteComment(ctx *gin.Context) {
	var cDeleteDto dto.DeleteCommentDTO
	if err := c.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &cDeleteDto}).GetError(); err != nil {
		c.Logger.Error(err)
		c.Fail(&Response{Code: ErrCodeDeleteComment, Msg: err.Error()})
		return
	}
	err := c.Service.DeleteCommentById(ctx, &cDeleteDto)
	if err != nil {
		c.Logger.Error(err)
		c.Fail(&Response{Code: ErrCodeDeleteComment, Msg: err.Error()})
		return
	}
	c.Success(&Response{})
}
