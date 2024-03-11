package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
)

const (
	ErrCodeGetCommentList = iota + 20001
	ErrCodeAddComment
	ErrCodeDeleteComment
)

type CommentApi struct {
	BaseApi
	Service *service.CommentService
}

func NewCommentApi() CommentApi {
	return CommentApi{
		BaseApi: NewBaseApi(),
		Service: service.NewCommentService(),
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
		c.Fail(&Response{Code: ErrCodeGetCommentList, Msg: err.Error()})
		return
	}

	comments, total, err := c.Service.GetCommentList(ctx, &cListDto)
	if err != nil {
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
		c.Fail(&Response{Code: ErrCodeAddComment, Msg: err.Error()})
		return
	}

	comment, err := c.Service.AddComment(ctx, &cAddDto)
	if err != nil {
		c.Fail(&Response{Code: ErrCodeAddComment, Msg: err.Error()})
		return
	}

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
		c.Fail(&Response{Code: ErrCodeDeleteComment, Msg: err.Error()})
		return
	}
	err := c.Service.DeleteCommentById(ctx, &cDeleteDto)
	if err != nil {
		c.Fail(&Response{Code: ErrCodeDeleteComment, Msg: err.Error()})
		return
	}
	c.Success(&Response{})
}
