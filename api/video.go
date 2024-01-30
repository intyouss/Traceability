package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
	"github.com/jinzhu/copier"
)

const (
	ErrCodeGetVideoFeed = iota + 30001
	ErrCodeGetUserVideoList
)

type VideoApi struct {
	BaseApi
	UserApi
	Service *service.VideoService
}

func NewVideoApi() VideoApi {
	return VideoApi{
		BaseApi: NewBaseApi(),
		UserApi: NewUserApi(),
		Service: service.NewVideoService(),
	}
}

// GetVideoFeed 获取视频feed流
// @Summary 获取视频feed流
// @Description 获取视频feed流
// @Param page formData int false "页码"
// @Param limit formData int false "每页数量"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/public/video/feed [post]
func (v *VideoApi) GetVideoFeed(ctx *gin.Context) {
	// 绑定并验证参数
	var vListDTO dto.VideoListDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &vListDTO}).GetError(); err != nil {
		v.Fail(&Response{Msg: err.Error()})
		return
	}

	// 调用service
	videos, total, err := v.Service.GetVideoList(ctx, &vListDTO)
	if err != nil {
		v.ServerError(&Response{Code: ErrCodeGetVideoFeed, Msg: err.Error()})
		return
	}
	if total == 0 {
		v.Success(&Response{
			Data:  []*dto.Video{},
			Total: total,
		})
		return
	}
	var authorMap map[uint]*models.User
	for _, video := range videos {
		authorMap[video.AuthorID] = nil
	}
	var authorIds []uint
	for authorId := range authorMap {
		authorIds = append(authorIds, authorId)
	}

	// 调用userApi
	authors, err := v.UserApi.Service.GetUserListByIds(ctx, authorIds)
	if err != nil {
		v.ServerError(&Response{Code: ErrCodeGetVideoFeed, Msg: err.Error()})
		return
	}

	for _, author := range authors {
		authorMap[author.ID] = author
	}

	// 组装数据
	var videoList = make([]*dto.Video, len(videos))
	_ = copier.Copy(&videoList, &videos)

	for i, video := range videoList {
		var user = new(dto.User)
		_ = copier.Copy(user, authorMap[videos[i].ID])
		video.Author = user
	}

	v.Success(&Response{
		Data:  videoList,
		Total: total,
	})
}

// GetUserVideoList 获取用户发布视频列表
// @Summary 获取用户发布视频列表
// @Description 获取用户发布视频列表
// @Param user_id query int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/list [get]
func (v *VideoApi) GetUserVideoList(ctx *gin.Context) {
	var idDTO dto.CommonUserIDDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &idDTO}).GetError(); err != nil {
		v.Fail(&Response{Msg: err.Error()})
		return
	}

	videosDao, err := v.Service.GetVideoListByUserId(ctx, &idDTO)
	if err != nil {
		v.ServerError(&Response{Code: ErrCodeGetUserVideoList, Msg: err.Error()})
		return
	}
	if len(videosDao) == 0 {
		v.Success(&Response{
			Data: []*dto.Video{},
		})
		return
	}

	var videos = make([]*dto.Video, len(videosDao))
	_ = copier.Copy(&videos, &videosDao)

	var userDao *models.User
	userDao, err = v.UserApi.Service.GetUserById(ctx, &idDTO)
	if err != nil {
		v.ServerError(&Response{Code: ErrCodeGetUserVideoList, Msg: err.Error()})
		return
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)

	var videoList []*dto.Video
	for _, video := range videos {
		video.Author = user
		videoList = append(videoList, video)
	}

	v.Success(&Response{
		Data: videoList,
	})
}

// PublishVideo 发布视频
// @Summary 发布视频
// @Description 发布视频
// @Param id formData int true "用户id"
// @Param title formData string true "视频标题"
// @Param data formData file true "视频文件"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/publish [post]
func (v *VideoApi) PublishVideo(ctx *gin.Context) {

}

// DeleteVideo 删除视频
// @Summary 删除视频
// @Description 删除视频
// @Param id formData int true "用户id"
// @Param video_id formData int true "视频id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/delete [delete]
func (v *VideoApi) DeleteVideo(ctx *gin.Context) {

}
