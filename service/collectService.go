package service

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var CollectServiceIns *CollectService

type CollectService struct {
	BaseService
	UserDao  *dao.UserDao
	VideoDao *dao.VideoDao
	Dao      *dao.CollectDao
}

func NewCollectService() *CollectService {
	if CollectServiceIns == nil {
		CollectServiceIns = &CollectService{
			Dao:      dao.NewCollectDao(),
			UserDao:  dao.NewUserDao(),
			VideoDao: dao.NewVideoDao(),
		}
	}
	return CollectServiceIns
}

// GetCollectList 获取收藏列表
func (l *CollectService) GetCollectList(ctx context.Context, cListDto *dto.CollectListDTO) ([]*dto.Video, error) {
	if !l.UserDao.IsExist(ctx, cListDto.UserID) {
		return nil, errors.New("user not exist")
	}
	myUserID := ctx.Value(global.LoginUser).(models.LoginUser).ID
	if myUserID != cListDto.UserID {
		return nil, errors.New("don't have permission")
	}
	collectListDao, err := l.Dao.GetCollectListByUserId(ctx, cListDto.UserID)
	if err != nil {
		return nil, err
	}
	if len(collectListDao) == 0 {
		return nil, nil
	}
	collectVideoIdList := make([]uint, 0, len(collectListDao))
	for _, collect := range collectListDao {
		collectVideoIdList = append(collectVideoIdList, collect.VideoID)
	}

	collectVideoListDao, err := l.VideoDao.GetVideoListByVideoId(ctx, collectVideoIdList)
	if err != nil {
		return nil, err
	}

	collectVideoUserIdMap := make(map[uint]*models.User, len(collectVideoIdList))
	for _, collect := range collectVideoListDao {
		collectVideoUserIdMap[collect.AuthorID] = nil
	}

	collectVideoUserIdList := make([]uint, 0, len(collectVideoUserIdMap))
	for userId := range collectVideoUserIdMap {
		collectVideoUserIdList = append(collectVideoUserIdList, userId)
	}

	collectUserList, err := l.UserDao.GetUserListByIds(ctx, collectVideoUserIdList)
	if err != nil {
		return nil, err
	}

	for _, user := range collectUserList {
		collectVideoUserIdMap[user.ID] = user
	}

	var collectVideoList []*dto.Video
	for _, video := range collectVideoListDao {
		var collectVideo = new(dto.Video)
		_ = copier.Copy(collectVideo, video)
		var user = new(dto.User)
		_ = copier.Copy(user, collectVideoUserIdMap[video.AuthorID])
		collectVideo.Author = user
		collectVideoList = append(collectVideoList, collectVideo)
	}
	return collectVideoList, nil
}

// CollectAction 收藏操作
func (l *CollectService) CollectAction(ctx context.Context, collectActionDto *dto.CollectActionDTO) error {
	if !l.VideoDao.IsExist(ctx, collectActionDto.VideoID) {
		return errors.New("video not exist")
	}
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	isCollected, err := l.Dao.IsCollected(ctx, collectActionDto.VideoID)
	if err != nil {
		return err
	}
	switch collectActionDto.ActionType {
	case 1:
		if isCollected {
			return errors.New("already collected")
		}
		return l.Dao.CreateCollectTransaction(ctx, userId, collectActionDto.VideoID)
	case 2:
		if !isCollected {
			return errors.New("do not have this collect relation")
		}
		return l.Dao.DeleteCollectTransaction(ctx, userId, collectActionDto.VideoID)
	default:
		return errors.New("action type error")
	}
}
