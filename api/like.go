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
	UserApi  UserApi
	VideoApi VideoApi
	Service  *service.LikeService
}

func NewLikeApi() LikeApi {
	return LikeApi{
		BaseApi:  NewBaseApi(),
		UserApi:  NewUserApi(),
		VideoApi: NewVideoApi(),
		Service:  service.NewLikeService(),
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
		l.Logger.Error(err)
		l.Fail(&Response{Code: ErrCodeGetLikeList, Msg: err.Error()})
		return
	}

	if !l.UserApi.Service.IsExist(ctx, kListDto.UserID) {
		l.Fail(&Response{Code: ErrCodeGetLikeList, Msg: "user not exist"})
		return
	}

	likeListDao, err := l.Service.GetLikeList(ctx, &kListDto)
	if err != nil {
		l.Logger.Error(err)
		l.Fail(&Response{Code: ErrCodeGetLikeList, Msg: err.Error()})
		return
	}
	if len(likeListDao) != 0 {
		likeVideoIdList := make([]uint, 0, len(likeListDao))
		for _, like := range likeListDao {
			likeVideoIdList = append(likeVideoIdList, like.VideoID)
		}

		likeVideoListDao, err := l.VideoApi.Service.GetVideoListByVideoId(ctx, likeVideoIdList)
		if err != nil {
			l.Logger.Error(err)
			l.Fail(&Response{Code: ErrCodeGetLikeList, Msg: err.Error()})
			return
		}

		likeVideoUserIdMap := make(map[uint]*models.User, len(likeVideoIdList))
		for _, like := range likeVideoListDao {
			likeVideoUserIdMap[like.AuthorID] = nil
		}

		likeVideoUserIdList := make([]uint, 0, len(likeVideoUserIdMap))
		for userId := range likeVideoUserIdMap {
			likeVideoUserIdList = append(likeVideoUserIdList, userId)
		}

		likeUserList, err := l.UserApi.Service.GetUserListByIds(ctx, likeVideoUserIdList)
		if err != nil {
			l.Logger.Error(err)
			l.Fail(&Response{Code: ErrCodeGetLikeList, Msg: err.Error()})
			return
		}

		for _, user := range likeUserList {
			likeVideoUserIdMap[user.ID] = user
		}

		var likeVideoList []*dto.Video
		for _, video := range likeVideoListDao {
			var likeVideo = new(dto.Video)
			_ = copier.Copy(likeVideo, video)
			var user = new(dto.User)
			_ = copier.Copy(user, likeVideoUserIdMap[video.AuthorID])
			likeVideo.Author = user
			likeVideoList = append(likeVideoList, likeVideo)
		}

		l.Success(&Response{
			Data: gin.H{
				"videos": likeVideoList,
			}})
		return
	}
	l.Success(&Response{
		Data: gin.H{
			"videos": []*dto.Video{},
		}})
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
		l.Logger.Error(err)
		l.Fail(&Response{Code: ErrCodeLikeAction, Msg: err.Error()})
		return
	}

	if !l.VideoApi.Service.IsExist(ctx, likeActionDto.VideoID) {
		l.Fail(&Response{Code: ErrCodeLikeAction, Msg: "video not exist"})
		return
	}

	if err := l.Service.LikeAction(ctx, &likeActionDto); err != nil {
		l.Logger.Error(err)
		l.Fail(&Response{Code: ErrCodeLikeAction, Msg: err.Error()})
		return
	}
	l.Success(&Response{})
}
