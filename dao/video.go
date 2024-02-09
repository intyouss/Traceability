package dao

import (
	"bytes"
	"context"
	"errors"
	"time"

	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
	"github.com/intyouss/Traceability/utils"
	"github.com/minio/minio-go/v7"
)

var VideoDaoIns *VideoDao

type VideoDao struct {
	*BaseDao
	OSS *utils.MinioClient
}

func NewVideoDao() *VideoDao {
	if VideoDaoIns == nil {
		VideoDaoIns = &VideoDao{
			BaseDao: NewBaseDao(),
			OSS:     global.OSS,
		}
	}
	return VideoDaoIns
}

// GetVideoList 获取视频列表
func (v *VideoDao) GetVideoList(ctx context.Context, vListDTO *dto.VideoListDTO) ([]*models.Video, int64, error) {
	var videos []*models.Video
	var total int64
	err := v.DB.Model(&models.Video{}).WithContext(ctx).
		Scopes(Paginate(vListDTO.CommonPageDTO)).Find(&videos).
		Offset(-1).Limit(-1).Count(&total).Error
	return videos, total, err
}

// GetVideoListByUserId 根据用户id获取视频列表
func (v *VideoDao) GetVideoListByUserId(ctx context.Context, idDTO *dto.CommonUserIDDTO) ([]*models.Video, error) {
	var videos []*models.Video
	err := v.DB.Model(&models.Video{}).WithContext(ctx).Where("user_id = ?", idDTO.ID).Find(&videos).Error
	return videos, err
}

// SaveVideoInfo 保存视频信息
func (v *VideoDao) SaveVideoInfo(ctx context.Context, upload *dto.VideoPublishDTO) error {
	playUrl, coverUrl, err := v.GetRemoteVideoInfo(ctx, upload)
	if err != nil {
		return err
	}
	video := &models.Video{
		AuthorID:     ctx.Value(global.LoginUser).(models.LoginUser).ID,
		Title:        upload.Title,
		PlayUrl:      playUrl,
		CoverUrl:     coverUrl,
		LikeCount:    0,
		CommentCount: 0,
	}
	return v.DB.WithContext(ctx).Create(video).Error
}

// GetRemoteVideoInfo 获取远程视频及封面url
func (v *VideoDao) GetRemoteVideoInfo(ctx context.Context, upload *dto.VideoPublishDTO) (playURL, coverURL string, err error) {
	hours, days := 24, 7
	urls, err := v.OSS.GetFileURL(
		ctx, "oss", "videos/"+upload.Title+".mp4", time.Hour*time.Duration(hours*days))
	if err != nil {
		return "", "", err
	}
	playURL = urls.String()
	urls, err = v.OSS.GetFileURL(
		ctx, "oss", "images/"+upload.Title+".png", time.Hour*time.Duration(hours*days))
	if err != nil {
		return "", "", err
	}
	coverURL = urls.String()
	return
}

// UploadVideo 上传视频
func (v *VideoDao) UploadVideo(ctx context.Context, upload *dto.VideoPublishDTO) error {
	title := upload.Title
	reader := bytes.NewReader(upload.Data)
	return v.OSS.UploadSizeFile(
		ctx, "oss", "videos/"+title+".mp4", reader, reader.Size(), minio.PutObjectOptions{
			ContentType: "video/mp4",
		},
	)
}

// UploadCoverImage 上传封面
func (v *VideoDao) UploadCoverImage(ctx context.Context, upload *dto.VideoPublishDTO) error {
	title := upload.Title
	coverBytes := bytes.NewReader(upload.CoverImageData)
	return v.OSS.UploadSizeFile(
		ctx, "oss", "images/"+title+".png", coverBytes, coverBytes.Size(), minio.PutObjectOptions{
			ContentType: "image/png",
		},
	)
}

// DeleteVideo 删除视频
func (v *VideoDao) DeleteVideo(ctx context.Context, deleteDTO *dto.VideoDeleteDTO) (*models.Video, error) {
	var video models.Video
	err := v.DB.WithContext(ctx).Model(&models.Video{}).First(&video, deleteDTO.VideoID).Error
	if err != nil {
		return nil, errors.New("video not found")
	}
	if video.AuthorID != ctx.Value(global.LoginUser).(models.LoginUser).ID {
		return nil, errors.New("no permission to delete this video")
	}
	err = v.DB.WithContext(ctx).Unscoped().Delete(&models.Video{}, deleteDTO.VideoID).Error
	return &video, err
}

// DeleteRemoteVideo 删除远程视频
func (v *VideoDao) DeleteRemoteVideo(ctx context.Context, video *models.Video) error {
	videoName := video.Title + ".mp4"
	return v.OSS.RemoveFile(ctx, "oss", "videos/"+videoName)
}

// DeleteRemoteCoverImage 删除远程封面
func (v *VideoDao) DeleteRemoteCoverImage(ctx context.Context, video *models.Video) error {
	imageName := video.Title + ".png"
	return v.OSS.RemoveFile(ctx, "oss", "images/"+imageName)
}
