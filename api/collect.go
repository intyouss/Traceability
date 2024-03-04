package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
	"github.com/jinzhu/copier"
)

const (
	ErrCodeGetCollectList = iota + 70001
	ErrCodeCollectAction
)

type CollectApi struct {
	BaseApi
	UserApi  UserApi
	VideoApi VideoApi
	Service  *service.CollectService
}

func NewCollectApi() CollectApi {
	return CollectApi{
		BaseApi:  NewBaseApi(),
		UserApi:  NewUserApi(),
		VideoApi: NewVideoApi(),
		Service:  service.NewCollectService(),
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

	if !l.UserApi.Service.IsExist(ctx, cListDto.UserID) {
		l.Fail(&Response{Code: ErrCodeGetCollectList, Msg: "user not exist"})
		return
	}

	collectListDao, err := l.Service.GetCollectList(ctx, &cListDto)
	if err != nil {
		l.Fail(&Response{Code: ErrCodeGetCollectList, Msg: err.Error()})
		return
	}
	if len(collectListDao) != 0 {
		collectVideoIdList := make([]uint, 0, len(collectListDao))
		for _, collect := range collectListDao {
			collectVideoIdList = append(collectVideoIdList, collect.VideoID)
		}

		collectVideoListDao, err := l.VideoApi.Service.GetVideoListByVideoId(ctx, collectVideoIdList)
		if err != nil {
			l.Fail(&Response{Code: ErrCodeGetCollectList, Msg: err.Error()})
			return
		}

		collectVideoUserIdMap := make(map[uint]*models.User, len(collectVideoIdList))
		for _, collect := range collectVideoListDao {
			collectVideoUserIdMap[collect.AuthorID] = nil
		}

		collectVideoUserIdList := make([]uint, 0, len(collectVideoUserIdMap))
		for userId := range collectVideoUserIdMap {
			collectVideoUserIdList = append(collectVideoUserIdList, userId)
		}

		collectUserList, err := l.UserApi.Service.GetUserListByIds(ctx, collectVideoUserIdList)
		if err != nil {
			l.Fail(&Response{Code: ErrCodeGetCollectList, Msg: err.Error()})
			return
		}

		for _, user := range collectUserList {
			collectVideoUserIdMap[user.ID] = user
		}

		var collectVideoList []*dto.Video
		for _, video := range collectVideoListDao {
			var collectVideo = new(dto.Video)
			_ = copier.Copy(collectVideo, video)
			var user = new(dto.User)
			_ = copier.Copy(user, collectVideoUserIdMap[video.AuthorID])
			collectVideo.Author = user
			collectVideoList = append(collectVideoList, collectVideo)
		}

		l.Success(&Response{
			Data: gin.H{
				"videos": collectVideoList,
			}})
		return
	}
	l.Success(&Response{
		Data: gin.H{
			"videos": []*dto.Video{},
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

	if !l.VideoApi.Service.IsExist(ctx, collectActionDto.VideoID) {
		l.Fail(&Response{Code: ErrCodeCollectAction, Msg: "video not exist"})
		return
	}

	if err := l.Service.CollectAction(ctx, &collectActionDto); err != nil {
		l.Fail(&Response{Code: ErrCodeCollectAction, Msg: err.Error()})
		return
	}
	l.Success(&Response{})
}
