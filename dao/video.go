package dao

import (
	"context"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var VideoDaoIns *VideoDao

type VideoDao struct {
	*BaseDao
}

func NewVideoDao() *VideoDao {
	if VideoDaoIns == nil {
		VideoDaoIns = &VideoDao{
			BaseDao: NewBaseDao(),
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
