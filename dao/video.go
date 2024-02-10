package dao

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
	"github.com/intyouss/Traceability/utils"
	"github.com/minio/minio-go/v7"
)

const (
	VideoBucket = "oss"
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

// IsExist 判断视频是否存在
func (v *VideoDao) IsExist(ctx context.Context, videoId uint) bool {
	err := v.DB.Model(&models.Video{}).WithContext(ctx).First(&models.Video{}, videoId).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// GetVideoList 获取视频列表
func (v *VideoDao) GetVideoList(
	ctx context.Context, vListDTO *dto.VideoListDTO,
) (videos []*models.Video, nextTime int64, err error) {
	var latestTime time.Time
	if vListDTO.LatestTime != nil && *vListDTO.LatestTime == 0 {
		latestTime = time.Now()
		err = v.DB.Model(&models.Video{}).WithContext(ctx).Where("created_at <= ?", latestTime).
			Order("id DESC").Find(&videos).Error
		nextTime = videos[0].CreatedAt.UnixMilli()
		return videos, nextTime, err
	}
	latestTime = time.UnixMilli(*vListDTO.LatestTime)
	err = v.DB.Model(&models.Video{}).WithContext(ctx).Where("created_at > ?", latestTime).
		Order("id DESC").Find(&videos).Error
	if len(videos) != 0 {
		nextTime = videos[0].CreatedAt.UnixMilli()
	}
	return videos, nextTime, err
}

// GetVideoListByUserId 根据用户id获取视频列表
func (v *VideoDao) GetVideoListByUserId(ctx context.Context, idDTO *dto.CommonUserIDDTO) ([]*models.Video, error) {
	var videos []*models.Video
	err := v.DB.Model(&models.Video{}).WithContext(ctx).Where("author_id = ?", idDTO.ID).
		Order("id DESC").Find(&videos).Error
	return videos, err
}

// GetVideoListByVideoId 根据视频id列表获取视频列表
func (v *VideoDao) GetVideoListByVideoId(ctx context.Context, videoIds []uint) ([]*models.Video, error) {
	var videos []*models.Video
	err := v.DB.Model(&models.Video{}).WithContext(ctx).Where("id IN ?", videoIds).
		Find(&videos).Error
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
	err = v.DB.WithContext(ctx).Create(video).Error
	if err != nil {
		// 保存失败，删除远程视频和封面,回滚
		go func() {
			err = v.DeleteRemoteVideo(ctx, video)
			if err != nil {
				v.logger.Error(err)
			}
		}()
		go func() {
			err = v.DeleteRemoteCoverImage(ctx, video)
			if err != nil {
				v.logger.Error(err)
			}
		}()
		return err
	}
	return nil
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
	// 读取视频文件
	videoSize := upload.Data.Size
	videoData, err := upload.Data.Open()
	if err != nil {
		return err
	}
	defer videoData.Close()
	// 上传视频
	title := upload.Title
	fileName := "videos/" + title + ".mp4"
	err = v.OSS.UploadSizeFile(
		ctx, VideoBucket, fileName, videoData, videoSize, minio.PutObjectOptions{
			ContentType: "video/mp4",
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// UploadCoverImage 上传封面
func (v *VideoDao) UploadCoverImage(ctx context.Context, upload *dto.VideoPublishDTO) error {
	// 读取封面图片
	imageSize := upload.CoverImageData.Size
	imageData, err := upload.CoverImageData.Open()
	if err != nil {
		return err
	}
	defer imageData.Close()
	// 上传封面
	title := upload.Title
	fileName := "images/" + title + ".png"
	err = v.OSS.UploadSizeFile(
		ctx, VideoBucket, fileName, imageData, imageSize, minio.PutObjectOptions{
			ContentType: "image/png",
		},
	)
	if err != nil {
		return err
	}
	return nil
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
	videoName := "videos/" + video.Title + ".mp4"
	return v.OSS.RemoveFile(ctx, VideoBucket, videoName)
}

// DeleteRemoteCoverImage 删除远程封面
func (v *VideoDao) DeleteRemoteCoverImage(ctx context.Context, video *models.Video) error {
	imageName := "images/" + video.Title + ".png"
	return v.OSS.RemoveFile(ctx, VideoBucket, imageName)
}
