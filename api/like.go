package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
	"github.com/jinzhu/copier"
)

const (
	ErrCodeGetLikeList = iota + 40001
	ErrCodeLikeAction
)

type LikeApi struct {
	BaseApi
	UserApi
	VideoApi
	Service *service.LikeService
}

func NewLikeApi() LikeApi {
	return LikeApi{
		BaseApi: NewBaseApi(),
		Service: service.NewLikeService(),
	}
}

// GetLikeList 获取点赞列表, 用户只能查看自己的喜爱列表
// @Summary 获取用户点赞列表
// @Description 获取用户点赞列表
// @Param token header string true "token"
// @Param user_id query int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/like/list [get]
func (l LikeApi) GetLikeList(ctx *gin.Context) {
	var kListDto dto.LikeListDTO
	if err := l.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &kListDto}).GetError(); err != nil {
		l.Fail(&Response{Code: ErrCodeGetLikeList, Msg: err.Error()})
		return
	}

	likeListDao, err := l.Service.GetLikeList(ctx, &kListDto)
	if err != nil {
		l.Fail(&Response{Code: ErrCodeGetLikeList, Msg: err.Error()})
		return
	}
	if len(likeListDao) != 0 {
		idDTO := dto.CommonUserIDDTO{
			ID: likeListDao[0].UserID,
		}

		var likeVideoListDao []*models.Video
		likeVideoListDao, err = l.VideoApi.Service.GetVideoListByUserId(ctx, &idDTO)
		if err != nil {
			l.Fail(&Response{Code: ErrCodeGetLikeList, Msg: err.Error()})
			return
		}

		var likeVideoList []*dto.Video
		_ = copier.Copy(&likeVideoList, &likeVideoListDao)

		l.Success(&Response{Data: likeVideoList})
	}
	l.Success(&Response{Data: []*dto.Video{}})
}

// LikeAction 用户喜爱操作
// @Summary 用户喜爱操作
// @Description 用户喜爱操作
// @Param token header string true "token"
// @Param video_id formData int true "视频id"
// @Param action_type formData int true "操作类型 1:喜爱 2:取消喜爱"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/like/action [post]
func (l LikeApi) LikeAction(ctx *gin.Context) {
	var likeActionDto dto.LikeActionDTO
	if err := l.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &likeActionDto}).GetError(); err != nil {
		l.Fail(&Response{Code: ErrCodeLikeAction, Msg: err.Error()})
		return
	}

	if err := l.Service.LikeAction(ctx, &likeActionDto); err != nil {
		l.Fail(&Response{Code: ErrCodeLikeAction, Msg: err.Error()})
		return
	}
	l.Success(&Response{})
}
