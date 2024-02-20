package api

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
	"github.com/jinzhu/copier"
)

const (
	ErrCodeGetVideoFeed = iota + 30001
	ErrCodeGetUserVideoList
	ErrCodePublishVideo
	ErrCodeDeleteVideo
)

type VideoApi struct {
	BaseApi
	UserApi UserApi
	Service *service.VideoService
}

func NewVideoApi() VideoApi {
	return VideoApi{
		BaseApi: NewBaseApi(),
		UserApi: NewUserApi(),
		Service: service.NewVideoService(),
	}
}

// GetVideoFeed 获取首页视频feed流
// @Summary 获取视频feed流
// @Description 获取视频feed流
// @Param latest_time formData int true "最新时间"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/public/video/feed/index [get]
func (v VideoApi) GetVideoFeed(ctx *gin.Context) {
	// 绑定并验证参数
	var vListDTO dto.VideoListDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &vListDTO}).GetError(); err != nil {
		v.Logger.Error(err)
		v.Fail(&Response{Code: ErrCodeGetVideoFeed, Msg: err.Error()})
		return
	}

	// 调用service
	videos, nextTime, err := v.Service.GetVideoList(ctx, &vListDTO)
	if err != nil {
		v.Logger.Error(err)
		v.Fail(&Response{Code: ErrCodeGetVideoFeed, Msg: err.Error()})
		return
	}
	if len(videos) == 0 {
		v.Success(&Response{
			Data: gin.H{
				"videos":    []*dto.Video{},
				"next_time": nextTime,
			},
		})
		return
	}
	authorMap := make(map[uint]*models.User)
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
		v.Logger.Error(err)
		v.Fail(&Response{Code: ErrCodeGetVideoFeed, Msg: err.Error()})
		return
	}
	for _, author := range authors {
		authorMap[author.ID] = author
	}

	// 组装数据
	var videoList = make([]*dto.Video, 0, len(videos))
	for _, video := range videos {
		var videoDTO = new(dto.Video)
		_ = copier.Copy(videoDTO, video)
		videoDTO.CreatedAt = timeFormat(video.CreatedAt)
		var user = new(dto.User)
		_ = copier.Copy(user, authorMap[video.AuthorID])
		videoDTO.Author = user
		videoList = append(videoList, videoDTO)
	}

	v.Success(&Response{
		Data: gin.H{
			"videos":    videoList,
			"next_time": nextTime,
		},
	})
}

// GetUserVideoList 获取用户发布视频列表
// @Summary 获取用户发布视频列表
// @Description 获取用户发布视频列表
// @Param user_id query int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/default/video/list [get]
func (v VideoApi) GetUserVideoList(ctx *gin.Context) {
	var idDTO dto.CommonUserIDDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &idDTO}).GetError(); err != nil {
		v.Logger.Error(err)
		v.Fail(&Response{Code: ErrCodeGetUserVideoList, Msg: err.Error()})
		return
	}

	if !v.UserApi.Service.IsExist(ctx, uint(idDTO.ID)) {
		v.Fail(&Response{Code: ErrCodeGetUserVideoList, Msg: "user not exist"})
		return
	}

	videosDao, err := v.Service.GetVideoListByUserId(ctx, &idDTO)
	if err != nil {
		v.Logger.Error(err)
		v.Fail(&Response{Code: ErrCodeGetUserVideoList, Msg: err.Error()})
		return
	}
	if len(videosDao) == 0 {
		v.Success(&Response{
			Data: gin.H{
				"videos": []*dto.Video{},
			},
		})
		return
	}

	var userDao *models.User
	userDao, err = v.UserApi.Service.GetUserById(ctx, &idDTO)
	if err != nil {
		v.Logger.Error(err)
		v.Fail(&Response{Code: ErrCodeGetUserVideoList, Msg: err.Error()})
		return
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)

	var videos = make([]*dto.Video, len(videosDao))
	for _, video := range videosDao {
		var videoDTO = new(dto.Video)
		_ = copier.Copy(videoDTO, video)
		videoDTO.CreatedAt = timeFormat(video.CreatedAt)
		videoDTO.Author = user
		videos = append(videos, videoDTO)
	}

	v.Success(&Response{
		Data: gin.H{
			"videos": videos,
		},
	})
}

// PublishVideo 发布视频
// @Summary 发布视频
// @Description 发布视频
// @Param token header string true "token"
// @Param title formData string true "视频标题"
// @Param cover_image_data formData file true "封面图片"
// @Param data formData file true "视频文件"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/publish [post]
// PS: 视频与图片文件的参数校验失效，需要手动处理
func (v VideoApi) PublishVideo(ctx *gin.Context) {
	var videoPublishDTO dto.VideoPublishDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &videoPublishDTO}).GetError(); err != nil {
		v.Logger.Error(err)
		v.Fail(&Response{Code: ErrCodePublishVideo, Msg: err.Error()})
		return
	}

	err := v.Service.PublishVideo(ctx, &videoPublishDTO)
	if err != nil {
		v.Logger.Error(err)
		v.Fail(&Response{Code: ErrCodePublishVideo, Msg: err.Error()})
		return
	}

	v.Success(&Response{})
}

// DeleteVideo 删除视频
// @Summary 删除视频
// @Description 删除视频
// @Param token header string true "token"
// @Param video_id formData int true "视频id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/delete [delete]
func (v VideoApi) DeleteVideo(ctx *gin.Context) {
	var videoDTO dto.VideoDeleteDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &videoDTO}).GetError(); err != nil {
		v.Logger.Error(err)
		v.Fail(&Response{Code: ErrCodeDeleteVideo, Msg: err.Error()})
		return
	}

	err := v.Service.DeleteVideo(ctx, &videoDTO)
	if err != nil {
		v.Logger.Error(err)
		v.Fail(&Response{Code: ErrCodeDeleteVideo, Msg: err.Error()})
		return
	}

	v.Success(&Response{})
}

// timeFormat 时间格式化
func timeFormat(t time.Time) string {
	since := time.Since(t)
	switch {
	case since < time.Minute: // 如果是一分钟内的时间，返回刚刚
		return "刚刚"
	case since < time.Hour: // 如果是一小时内的时间，返回是几分钟前
		return strings.Split(since.String(), "m")[0] + "分钟前"
	case since < 24*time.Hour: // 如果是超过一个小时的时间，返回是几小时前
		return strings.Split(since.String(), "h")[0] + "小时前"
	case since < 7*24*time.Hour: // 如果超过一天但是在一周内，返回是几天前
		x, _ := strconv.Atoi(strings.Split(since.String(), "h")[0])
		return strconv.Itoa(x/24) + "天前"
	case since < 21*24*time.Hour: // 如果超过一周，但不超过三周，返回是几周前
		x, _ := strconv.Atoi(strings.Split(since.String(), "h")[0])
		return strconv.Itoa(x/(7*24)) + "周前"
	default: // 如果超过三周，返回年月日
		return t.Format("2006-01-02")
	}
}
