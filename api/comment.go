package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
)

const (
	ErrCodeGetCommentList = iota + 20001
	ErrCodeCommentAction
)

type CommentApi struct {
	BaseApi
	UserApi
	Service *service.CommentService
}

func NewCommentApi() CommentApi {
	return CommentApi{
		BaseApi: NewBaseApi(),
		UserApi: NewUserApi(),
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
// @Router /api/v1/public/comment/list [post]
func (c CommentApi) GetCommentList(ctx *gin.Context) {
	var cListDto dto.CommentListDTO
	if err := c.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &cListDto}).GetError(); err != nil {
		c.Fail(&Response{Msg: err.Error()})
		return
	}

	comments, total, err := c.Service.GetCommentList(&cListDto)
	if err != nil {
		c.ServerError(&Response{Code: ErrCodeGetCommentList, Msg: err.Error()})
		return
	}

	c.Success(&Response{
		Data:  comments,
		Total: total,
	})
}

// CommentAction 评论操作
// @Summary 评论操作
// @Description 评论操作
// @Param action_type formData int true "操作类型 1:添加 2:删除"
// @Param comment_id formData int false "评论id"
// @Param video_id formData int true "视频id"
// @Param content formData string false "评论内容"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/comment/action [post]
func (c CommentApi) CommentAction(ctx *gin.Context) {
	var cActionDto dto.CommentActionDTO
	if err := c.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &cActionDto}).GetError(); err != nil {
		c.Fail(&Response{Msg: err.Error()})
		return
	}
	userId := ctx.GetUint(global.LoginUser)
	cActionDto.UserID = userId
	comment, err := c.Service.CommentAction(&cActionDto)
	if err != nil {
		c.ServerError(&Response{Code: ErrCodeCommentAction, Msg: err.Error()})
		return
	}
	if cActionDto.ActionType == 1 {
		var IdDTO dto.CommonIDDTO
		IdDTO.ID = userId
		user, err := c.UserApi.Service.GetUserById(&IdDTO)
		if err != nil {
			c.ServerError(&Response{Code: ErrCodeGetUserById, Msg: err.Error()})
			return
		}
		c.Success(&Response{
			Data: gin.H{
				"comment": comment,
				"user":    user,
			},
		})
	} else {
		c.Success(&Response{})
	}
}
