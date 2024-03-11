package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
)

const (
	ErrCodeRelationAction = iota + 50001
	ErrCodeGetFocusList
	ErrCodeGetFansList
)

type RelationApi struct {
	BaseApi
	Service *service.RelationService
}

func NewRelationApi() RelationApi {
	return RelationApi{
		BaseApi: NewBaseApi(),
		Service: service.NewRelationService(),
	}
}

// RelationAction 关注/取消关注
// @Summary 关注/取消关注
// @Description 关注/取消关注
// @Param token header string true "token"
// @Param action_type formData int true "1:关注 2:取消关注"
// @Param user_id formData int true "想要关注/取关的用户ID"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/relation/action [post]
func (r RelationApi) RelationAction(ctx *gin.Context) {
	var relationDto dto.RelationActionDto
	if err := r.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &relationDto}).GetError(); err != nil {
		r.Fail(&Response{Code: ErrCodeRelationAction, Msg: err.Error()})
		return
	}

	err := r.Service.RelationAction(ctx, relationDto)
	if err != nil {
		r.Fail(&Response{Code: ErrCodeRelationAction, Msg: err.Error()})
		return
	}

	r.Success(&Response{})
}

// GetFocusList 关注列表
// @Summary 关注列表
// @Description 关注列表
// @Param token header string true "token"
// @Param user_id query int true "用户ID"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/relation/focus/list [get]
func (r RelationApi) GetFocusList(ctx *gin.Context) {
	var focusListDto dto.FocusListDto
	if err := r.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &focusListDto}).GetError(); err != nil {
		r.Fail(&Response{Code: ErrCodeGetFocusList, Msg: err.Error()})
		return
	}

	total, focusList, err := r.Service.GetFocusList(ctx, focusListDto)
	if err != nil {
		r.Fail(&Response{Code: ErrCodeGetFocusList, Msg: err.Error()})
		return
	}
	if total == 0 {
		r.Success(&Response{
			Data: gin.H{
				"users": []dto.User{},
			},
			Total: 0,
		})
		return
	}

	r.Success(&Response{
		Data: gin.H{
			"users": focusList,
		},
		Total: total,
	})
}

// GetFansList 粉丝列表
// @Summary 粉丝列表
// @Description 粉丝列表
// @Param token header string true "token"
// @Param user_id query int true "用户ID"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/relation/fans/list [get]
func (r RelationApi) GetFansList(ctx *gin.Context) {
	var fansListDto dto.FansListDto
	if err := r.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &fansListDto}).GetError(); err != nil {
		r.Fail(&Response{Code: ErrCodeGetFansList, Msg: err.Error()})
		return
	}

	total, fansList, err := r.Service.GetFansList(ctx, fansListDto)
	if err != nil {
		r.Fail(&Response{Code: ErrCodeGetFansList, Msg: err.Error()})
		return
	}
	if total == 0 {
		r.Success(&Response{
			Data: gin.H{
				"users": []dto.User{},
			},
			Total: 0,
		})
		return
	}

	r.Success(&Response{
		Data: gin.H{
			"users": fansList,
		},
		Total: total,
	})
}
