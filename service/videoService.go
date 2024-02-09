package service

import (
	"context"
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
func (v *VideoService) GetVideoList(ctx context.Context, vListDTO *dto.VideoListDTO) ([]*models.Video, int64, error) {
	return v.Dao.GetVideoList(ctx, vListDTO)
}

// GetVideoListByUserId 根据用户id获取视频列表
func (v *VideoService) GetVideoListByUserId(ctx context.Context, idDTO *dto.CommonUserIDDTO) ([]*models.Video, error) {
	return v.Dao.GetVideoListByUserId(ctx, idDTO)
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
