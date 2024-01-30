package service

import (
	"context"
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
