package dao

import (
	"context"
	"errors"
	"net/url"
	"strconv"
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
	UserDao     *UserDao
	RelationDao *RelationDao
	OSS         *utils.MinioClient
}

func NewVideoDao() *VideoDao {
	if VideoDaoIns == nil {
		VideoDaoIns = &VideoDao{
			BaseDao:     NewBaseDao(),
			UserDao:     NewUserDao(),
			RelationDao: NewRelationDao(),
			OSS:         global.OSS,
		}
	}
	return VideoDaoIns
}

// IsExist 判断视频是否存在
func (v *VideoDao) IsExist(ctx context.Context, videoId uint) bool {
	err := v.DB.Model(&models.Video{}).WithContext(ctx).First(&models.Video{}, videoId).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// GetAuthorIdByVideoId 根据视频id获取作者id
func (v *VideoDao) GetAuthorIdByVideoId(ctx context.Context, videoId uint) (uint, error) {
	var video models.Video
	err := v.DB.Model(&models.Video{}).WithContext(ctx).First(&video, videoId).Error
	return video.AuthorID, err
}

// GetVideoList 获取视频列表
func (v *VideoDao) GetVideoList(
	ctx context.Context, vListDTO *dto.VideoListDTO,
) (videos []*models.Video, nextTime string, err error) {
	//if vListDTO.LatestTime != nil && *vListDTO.LatestTime == 0 {
	//	latestTime = time.Now()
	//	err = v.DB.Model(&models.Video{}).WithContext(ctx).Where("created_at <= ?", latestTime).
	//		Order("id DESC").Find(&videos).Error
	//	nextTime = videos[0].CreatedAt.UnixMilli()
	//	return videos, nextTime, err
	//}
	latestTime, err := strconv.ParseInt(vListDTO.LatestTime, 10, 64)
	if err != nil {
		return nil, "", err
	}
	err = v.DB.Model(&models.Video{}).WithContext(ctx).Where("created_at > ?", time.UnixMilli(latestTime)).
		Order("id DESC").Find(&videos).Error
	if len(videos) == 0 {
		return nil, "", nil
	}
	nextTime = strconv.Itoa(int(videos[0].CreatedAt.UnixMilli()))
	return videos, nextTime, err
}

// GetFocusVideoList 获取关注视频列表
func (v *VideoDao) GetFocusVideoList(
	ctx context.Context, vListDTO *dto.VideoListDTO,
) (videoList []*models.Video, nextTime string, err error) {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	relations, err := v.RelationDao.GetFocusListByUserId(ctx, userId)
	if err != nil {
		return nil, "", err
	}
	if len(relations) == 0 {
		return nil, "", nil
	}
	var focusIds []uint
	for _, relation := range relations {
		focusIds = append(focusIds, relation.FocusID)
	}
	videos, _, err := v.GetVideoList(ctx, vListDTO)
	if err != nil {
		return nil, "", err
	}
	if len(videos) == 0 {
		return nil, "", nil
	}
	for _, video := range videos {
		for _, focusId := range focusIds {
			if video.AuthorID == focusId {
				videoList = append(videoList, video)
			}
		}
	}
	nextTime = strconv.Itoa(int(videoList[0].CreatedAt.UnixMilli()))
	return videoList, nextTime, nil
}

// GetFriendVideoList 获取好友视频列表
func (v *VideoDao) GetFriendVideoList(
	ctx context.Context, vListDTO *dto.VideoListDTO,
) (videoList []*models.Video, nextTime string, err error) {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	relations, err := v.RelationDao.GetFriendListByUserId(ctx, userId)
	if err != nil {
		return nil, "", err
	}
	if len(relations) == 0 {
		return nil, "", nil
	}
	var focusIds []uint
	for _, relation := range relations {
		focusIds = append(focusIds, relation.FocusID)
	}
	videos, _, err := v.GetVideoList(ctx, vListDTO)
	if err != nil {
		return nil, "", err
	}
	if len(videos) == 0 {
		return nil, "", nil
	}
	for _, video := range videos {
		for _, focusId := range focusIds {
			if video.AuthorID == focusId {
				videoList = append(videoList, video)
			}
		}
	}
	nextTime = strconv.Itoa(int(videoList[0].CreatedAt.UnixMilli()))
	return videoList, nextTime, nil
}

// GetVideoListByUserId 根据用户id获取视频列表
func (v *VideoDao) GetVideoListByUserId(ctx context.Context, idDTO *dto.CommonIDDTO) ([]*models.Video, error) {
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

// GetVideoSearchByTitle 根据标题模糊搜索视频
func (v *VideoDao) GetVideoSearchByTitle(ctx context.Context, searchDTO *dto.VideoSearchDTO) (videos []*models.Video, err error) {
	err = v.DB.Model(&models.Video{}).WithContext(ctx).
		Where("title LIKE ?", "%"+searchDTO.Key+"%").
		Find(&videos).Error
	return
}

// GetVideoSearchByAuthorAndTitle 根据标题与用户名模糊搜索视频
func (v *VideoDao) GetVideoSearchByAuthorAndTitle(
	ctx context.Context, searchDTO *dto.VideoSearchDTO,
) (videos []*models.Video, err error) {
	ids, err := v.UserDao.GetUserIdsBySearchKey(ctx, searchDTO.Key)
	if err != nil {
		return
	}
	err = v.DB.Model(&models.Video{}).WithContext(ctx).
		Where("title LIKE ?", "%"+searchDTO.Key+"%").
		Or("author_id IN ?", ids).Find(&videos).Error
	return
}

// SaveVideoInfo 保存视频信息
func (v *VideoDao) SaveVideoInfo(ctx context.Context, upload *dto.VideoPublishDTO) error {
	playUrl, coverUrl, err := v.GetRemoteVideoInfo(ctx, upload.Title)
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
	return v.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&video).Error; err != nil {
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
		// 更新视频数
		if err := v.UserDao.UpdateVideoCount(ctx, video.AuthorID, 1); err != nil {
			return err
		}
		return nil
	})
}

// GetRemoteVideoInfo 获取远程视频及封面url
func (v *VideoDao) GetRemoteVideoInfo(ctx context.Context, title string) (playURL, coverURL string, err error) {
	hours, days := 24, 7
	urls, err := v.OSS.GetFileURL(
		ctx, "oss", "videos/"+title+".mp4", time.Hour*time.Duration(hours*days))
	if err != nil {
		return "", "", err
	}
	playURL = urls.String()
	urls, err = v.OSS.GetFileURL(
		ctx, "oss", "images/"+title+".png", time.Hour*time.Duration(hours*days))
	if err != nil {
		return "", "", err
	}
	coverURL = urls.String()
	return
}

// CheckUrl 检查url是否过期
func (v *VideoDao) CheckUrl(accessUrl string) (bool, error) {
	parseUrl, err := url.Parse(accessUrl)
	if err != nil {
		return false, err
	}
	dateStr := parseUrl.Query().Get("X-Amz-Date")
	dateInt, err := time.Parse("20060102T150405Z", dateStr)
	if err != nil {
		return false, err
	}
	// 7天后过期,提前一个小时生成新的url
	hours, days := 24, 7
	now := time.Now().Add(time.Hour).UTC()
	return now.Before(dateInt.Add(time.Hour * time.Duration(hours*days))), nil
}

// UpdateUrl 检查视频列表所有url是否失效，更新url
func (v *VideoDao) UpdateUrl(ctx context.Context, videoList []*models.Video) error {
	for _, video := range videoList {
		firstOk, err := v.CheckUrl(video.PlayUrl)
		if err != nil {
			return err
		}
		secondOk, err := v.CheckUrl(video.CoverUrl)
		if err != nil {
			return err
		}
		if firstOk && secondOk {
			continue
		}
		video.PlayUrl, video.CoverUrl, err = v.GetRemoteVideoInfo(ctx, video.Title)
		if err != nil {
			return err
		}
		go func(vi *models.Video) {
			err = v.UpdateDBUrl(ctx, vi.ID, vi.PlayUrl, vi.CoverUrl)
			if err != nil {
				v.logger.Error(err)
			}
		}(video)
	}
	return nil
}

// UpdateDBUrl 更新数据库url
func (v *VideoDao) UpdateDBUrl(ctx context.Context, videoId uint, playUrl, coverUrl string) error {
	return v.DB.WithContext(ctx).Where("id = ?", videoId).
		Updates(&models.Video{PlayUrl: playUrl, CoverUrl: coverUrl}).Error
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
	err := v.DB.WithContext(ctx).Model(&models.Video{}).First(&video, deleteDTO.ID).Error
	if err != nil {
		return nil, errors.New("video not found")
	}
	if video.AuthorID != ctx.Value(global.LoginUser).(models.LoginUser).ID {
		return nil, errors.New("no permission to delete this video")
	}
	err = v.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = v.DB.WithContext(ctx).Unscoped().Delete(&models.Video{}, deleteDTO.ID).Error; err != nil {
			return err
		}
		// 更新视频数
		if err = v.UserDao.UpdateVideoCount(ctx, video.AuthorID, -1); err != nil {
			return err
		}
		return nil
	})
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

// UpdateCommentCount 更新评论数
func (v *VideoDao) UpdateCommentCount(ctx context.Context, videoId uint, count int) error {
	value := map[string]interface{}{"comment_count": gorm.Expr("comment_count + ?", count)}
	return v.DB.Model(&models.Video{}).WithContext(ctx).Where("id = ?", videoId).
		Updates(value).Error
}

// UpdateVideoLikeCount 更新点赞数
func (v *VideoDao) UpdateVideoLikeCount(ctx context.Context, videoId uint, count int) error {
	value := map[string]interface{}{"like_count": gorm.Expr("like_count + ?", count)}
	return v.DB.Model(&models.Video{}).WithContext(ctx).Where("id = ?", videoId).
		Updates(value).Error
}
