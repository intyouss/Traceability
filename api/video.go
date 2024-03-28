package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
)

const (
	ErrCodeGetVideoFeed = iota + 30001
	ErrCodeGetUserVideoList
	ErrCodeUploadVideo
	ErrCodeUploadImage
	ErrCodeSaveVideoInfo
	ErrCodeAbolishVideoUpload
	ErrCodeGetVideoInfo
	ErrCodeGetVideoSearch
)

type VideoApi struct {
	BaseApi
	Service *service.VideoService
}

func NewVideoApi() VideoApi {
	return VideoApi{
		BaseApi: NewBaseApi(),
		Service: service.NewVideoService(),
	}
}

// GetVideoFeed 获取视频feed流
// @Summary 获取视频feed流
// @Description 获取视频feed流
// @Param latest_time formData int true "最新时间"
// @Param type formData int true "feed类型"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/public/video/feed [get]

// GetVideoFeed 获取视频feed流 (auth)
// @Summary 获取视频feed流
// @Description 获取视频feed流
// @Param token header string true "token"
// @Param type formData int true "feed类型"
// @Param latest_time formData int true "最新时间"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/auth/video/feed [get]
func (v VideoApi) GetVideoFeed(ctx *gin.Context) {
	// 绑定并验证参数
	var vListDTO dto.VideoListDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &vListDTO}).GetError(); err != nil {

		v.Fail(&Response{Code: ErrCodeGetVideoFeed, Msg: err.Error()})
		return
	}

	// 调用service
	videos, nextTime, err := v.Service.GetVideoList(ctx, &vListDTO)
	if err != nil {
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

	v.Success(&Response{
		Data: gin.H{
			"videos":    videos,
			"next_time": nextTime,
		},
	})
}

// GetUserVideoList 获取其他用户发布视频列表
// @Summary 获取用户发布视频列表
// @Description 获取用户发布视频列表
// @Param user_id query int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/default/video/list [get]

// GetUserVideoList 获取用户发布视频列表
// @Summary 获取用户发布视频列表
// @Description 获取用户发布视频列表
// @Param token header string true "token"
// @Param user_id query int true "用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/list [get]
func (v VideoApi) GetUserVideoList(ctx *gin.Context) {
	var vListDTO dto.UserVideoListDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &vListDTO}).GetError(); err != nil {
		v.Fail(&Response{Code: ErrCodeGetUserVideoList, Msg: err.Error()})
		return
	}

	videos, err := v.Service.GetVideoListByUserId(ctx, &vListDTO)
	if err != nil {
		v.Fail(&Response{Code: ErrCodeGetUserVideoList, Msg: err.Error()})
		return
	}
	if len(videos) == 0 {
		v.Success(&Response{
			Data: gin.H{
				"videos": []*dto.Video{},
			},
		})
		return
	}

	v.Success(&Response{
		Data: gin.H{
			"videos": videos,
		},
	})
}

// GetVideoSearch 获取视频搜索结果
// @Summary 获取视频搜索结果
// @Description 获取视频搜索结果
// @Param key formData string true "搜索关键字"
// @Param type formData int true "搜索类型"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/public/video/search [get]
func (v VideoApi) GetVideoSearch(ctx *gin.Context) {
	var videoSearchDTO dto.VideoSearchDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &videoSearchDTO}).GetError(); err != nil {
		v.Fail(&Response{Code: ErrCodeGetVideoSearch, Msg: err.Error()})
		return
	}

	videos, err := v.Service.GetVideoSearch(ctx, &videoSearchDTO)
	if err != nil {
		v.Fail(&Response{Code: ErrCodeGetVideoSearch, Msg: err.Error()})
		return
	}
	if len(videos) == 0 {
		v.Success(&Response{
			Data: gin.H{
				"videos": []*dto.Video{},
			},
		})
		return
	}

	v.Success(&Response{
		Data: gin.H{
			"videos": videos,
		},
	})
}

// UploadVideo 上传视频
// @Summary 上传视频
// @Description 上传视频
// @Param token header string true "token"
// @Param data formData file true "视频文件"
// @Param title formData string true "标题"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/upload/video [post]
// PS: 视频与图片文件的参数校验失效，需要手动处理
func (v VideoApi) UploadVideo(ctx *gin.Context) {
	var videoUploadDTO dto.VideoUploadDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &videoUploadDTO}).GetError(); err != nil {
		v.Fail(&Response{Code: ErrCodeUploadVideo, Msg: err.Error()})
		return
	}

	playUrl, err := v.Service.UploadVideo(ctx, &videoUploadDTO)
	if err != nil {
		v.Fail(&Response{Code: ErrCodeUploadVideo, Msg: err.Error()})
		return
	}

	v.Success(&Response{
		Data: gin.H{
			"play_url": playUrl,
		},
	})
}

// UploadImage 上传视频封面
// @Summary 上传视频封面
// @Description 上传视频封面
// @Param token header string true "token"
// @Param cover_image_data formData file true "封面图片"
// @Param title formData string true "标题"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/upload/image [post]
// PS: 视频与图片文件的参数校验失效，需要手动处理
func (v VideoApi) UploadImage(ctx *gin.Context) {
	var imageUploadDTO dto.ImageUploadDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &imageUploadDTO}).GetError(); err != nil {
		v.Fail(&Response{Code: ErrCodeUploadImage, Msg: err.Error()})
		return
	}

	coverUrl, err := v.Service.UploadImage(ctx, &imageUploadDTO)
	if err != nil {
		v.Fail(&Response{Code: ErrCodeUploadImage, Msg: err.Error()})
		return
	}

	v.Success(&Response{
		Data: gin.H{
			"cover_image_url": coverUrl,
		},
	})
}

// SaveVideoInfo 保存视频
// @Summary 保存视频
// @Description 保存视频
// @Param token header string true "token"
// @Param title formData string true "标题"
// @Param video_url formData string true "视频地址"
// @Param cover_image_url formData string true "封面地址"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/publish [post]
func (v VideoApi) SaveVideoInfo(ctx *gin.Context) {
	var publishDTO dto.PublishDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &publishDTO}).GetError(); err != nil {
		v.Fail(&Response{Code: ErrCodeSaveVideoInfo, Msg: err.Error()})
		return
	}

	err := v.Service.SaveVideoInfo(ctx, &publishDTO)
	if err != nil {
		v.Fail(&Response{Code: ErrCodeSaveVideoInfo, Msg: err.Error()})
		return
	}

	v.Success(&Response{})
}

// AbolishVideoUpload 取消视频上传
// @Summary 取消视频上传
// @Description 取消视频上传
// @Param token header string true "token"
// @Param title formData string true "标题"
// @Param haveVideo formData bool true "是否有视频"
// @Param haveCoverImage formData bool true "是否有封面"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/upload/abolish [post]
func (v VideoApi) AbolishVideoUpload(ctx *gin.Context) {
	var abolishDTO dto.AbolishVideoUploadDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &abolishDTO}).GetError(); err != nil {
		v.Fail(&Response{Code: ErrCodeAbolishVideoUpload, Msg: err.Error()})
		return
	}

	err := v.Service.DeleteRemoteVideo(ctx, &abolishDTO)
	if err != nil {
		v.Fail(&Response{Code: ErrCodeAbolishVideoUpload, Msg: err.Error()})
		return
	}

	v.Success(&Response{})
}

// GetVideoInfo 获取视频信息
// @Summary 获取视频信息
// @Description 获取视频信息
// @Param token header string true "token"
// @Param id query int true "视频id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/video/info [get]
func (v VideoApi) GetVideoInfo(ctx *gin.Context) {
	var videoDTO dto.CommonIDDTO
	if err := v.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &videoDTO}).GetError(); err != nil {
		v.Fail(&Response{Code: ErrCodeGetVideoInfo, Msg: err.Error()})
		return
	}

	video, err := v.Service.GetVideoInfo(ctx, &videoDTO)
	if err != nil {
		v.Fail(&Response{Code: ErrCodeGetVideoInfo, Msg: err.Error()})
		return
	}

	v.Success(&Response{
		Data: gin.H{
			"video": video,
		},
	})
}
