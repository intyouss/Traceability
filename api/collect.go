package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
)

const (
	ErrCodeGetCollectList = iota + 70001
	ErrCodeCollectAction
)

type CollectApi struct {
	BaseApi
	Service *service.CollectService
}

func NewCollectApi() CollectApi {
	return CollectApi{
		BaseApi: NewBaseApi(),
		Service: service.NewCollectService(),
	}
}

// GetCollectList 获取收藏列表, 用户只能查看自己的收藏列表
// @Summary 获取用户收藏列表
// @Description 获取用户收藏列表
// @Param token header string true "token"
// @Param user_id query int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/collect/list [get]
func (l CollectApi) GetCollectList(ctx *gin.Context) {
	var cListDto dto.CollectListDTO
	if err := l.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &cListDto}).GetError(); err != nil {
		l.Fail(&Response{Code: ErrCodeGetCollectList, Msg: err.Error()})
		return
	}

	collectList, err := l.Service.GetCollectList(ctx, &cListDto)
	if err != nil {
		l.Fail(&Response{Code: ErrCodeGetCollectList, Msg: err.Error()})
		return
	}
	if len(collectList) == 0 {
		l.Success(&Response{
			Data: gin.H{
				"videos": []*dto.Video{},
			}})
		return
	}

	l.Success(&Response{
		Data: gin.H{
			"videos": collectList,
		}})
}

// CollectAction 用户收藏操作
// @Summary 用户收藏操作
// @Description 用户收藏操作
// @Param token header string true "token"
// @Param video_id formData int true "视频id"
// @Param action_type formData int true "操作类型 1:收藏 2:取消收藏"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/collect/action [post]
func (l CollectApi) CollectAction(ctx *gin.Context) {
	var collectActionDto dto.CollectActionDTO
	if err := l.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &collectActionDto}).GetError(); err != nil {
		l.Fail(&Response{Code: ErrCodeCollectAction, Msg: err.Error()})
		return
	}

	if err := l.Service.CollectAction(ctx, &collectActionDto); err != nil {
		l.Fail(&Response{Code: ErrCodeCollectAction, Msg: err.Error()})
		return
	}
	l.Success(&Response{})
}
