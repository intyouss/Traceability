package service

import (
	"context"
	"errors"
	"sync"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var VideoServiceIns *VideoService

type VideoService struct {
	BaseService
	Dao *dao.VideoDao
}

func NewVideoService() *VideoService {
	if VideoServiceIns == nil {
		VideoServiceIns = &VideoService{
			Dao: dao.NewVideoDao(),
		}
	}
	return VideoServiceIns
}

// GetVideoList 获取视频列表
func (v *VideoService) GetVideoList(
	ctx context.Context, vListDTO *dto.VideoListDTO,
) (videos []*models.Video, nextTime int64, err error) {
	switch vListDTO.Type {
	case 1:
		videos, nextTime, err = v.Dao.GetVideoList(ctx, vListDTO)
	case 2:
		videos, nextTime, err = v.Dao.GetFocusVideoList(ctx, vListDTO)
	case 3:
		videos, nextTime, err = v.Dao.GetFriendVideoList(ctx, vListDTO)
	case 4:
		videos, nextTime, err = v.Dao.GetVideoList(ctx, vListDTO)
	default:
		return nil, 0, errors.New("type error")
	}
	if err != nil {
		return nil, 0, err
	}
	if len(videos) == 0 {
		return nil, 0, nil
	}
	err = v.Dao.UpdateUrl(ctx, videos)
	if err != nil {
		return nil, 0, err
	}
	return videos, nextTime, nil
}

// GetVideoListByUserId 根据用户id获取视频列表
func (v *VideoService) GetVideoListByUserId(ctx context.Context, idDTO *dto.CommonUserIDDTO) ([]*models.Video, error) {
	videos, err := v.Dao.GetVideoListByUserId(ctx, idDTO)
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, nil
	}
	err = v.Dao.UpdateUrl(ctx, videos)
	return videos, err
}

// IsExist 判断视频是否存在
func (v *VideoService) IsExist(ctx context.Context, videoId uint) bool {
	return v.Dao.IsExist(ctx, videoId)
}

// GetVideoListByVideoId 根据视频id列表获取视频列表
func (v *VideoService) GetVideoListByVideoId(ctx context.Context, videoIds []uint) ([]*models.Video, error) {
	videos, err := v.Dao.GetVideoListByVideoId(ctx, videoIds)
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, nil
	}
	err = v.Dao.UpdateUrl(ctx, videos)
	return videos, err
}

func (v *VideoService) GetVideoSearch(
	ctx context.Context, searchDTO *dto.VideoSearchDTO,
) (videos []*models.Video, err error) {
	switch searchDTO.Type {
	case 1:
		videos, err = v.Dao.GetVideoSearchByAuthorAndTitle(ctx, searchDTO)
	case 2:
		videos, err = v.Dao.GetVideoSearchByTitle(ctx, searchDTO)
	}
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, nil
	}
	err = v.Dao.UpdateUrl(ctx, videos)
	return
}

// PublishVideo 发布视频
func (v *VideoService) PublishVideo(ctx context.Context, upload *dto.VideoPublishDTO) error {
	var wg sync.WaitGroup
	errChan := make(chan error)
	// 上传封面
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := v.Dao.UploadCoverImage(ctx, upload)
		if err != nil {
			errChan <- err
			return
		}
	}()
	// 上传视频
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := v.Dao.UploadVideo(ctx, upload)
		if err != nil {
			errChan <- err
			return
		}
	}()
	wg.Wait()
	select {
	case err := <-errChan:
		return err
	default:
		// 获取视频和封面的url
		err := v.Dao.SaveVideoInfo(ctx, upload)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteVideo 删除视频
func (v *VideoService) DeleteVideo(ctx context.Context, deleteDTO *dto.VideoDeleteDTO) error {
	video, err := v.Dao.DeleteVideo(ctx, deleteDTO)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	errChan := make(chan error)
	// 删除封面
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := v.Dao.DeleteRemoteCoverImage(ctx, video)
		if err != nil {
			errChan <- err
			return
		}
	}()
	// 删除视频
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := v.Dao.DeleteRemoteVideo(ctx, video)
		if err != nil {
			errChan <- err
			return
		}
	}()
	wg.Wait()
	select {
	case err = <-errChan:
		return err
	default:
		return nil
	}
}
