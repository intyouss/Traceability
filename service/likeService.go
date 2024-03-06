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

var LikeServiceIns *LikeService

type LikeService struct {
	BaseService
	UserDao  *dao.UserDao
	VideoDao *dao.VideoDao
	Dao      *dao.LikeDao
}

func NewLikeService() *LikeService {
	if LikeServiceIns == nil {
		LikeServiceIns = &LikeService{
			Dao:      dao.NewLikeDao(),
			UserDao:  dao.NewUserDao(),
			VideoDao: dao.NewVideoDao(),
		}
	}
	return LikeServiceIns
}

// GetLikeList 获取点赞列表
func (l *LikeService) GetLikeList(ctx context.Context, kListDto *dto.LikeListDTO) ([]*dto.Video, error) {
	if !l.UserDao.IsExist(ctx, kListDto.UserID) {
		return nil, errors.New("user not exist")
	}

	myUserID := ctx.Value(global.LoginUser).(models.LoginUser).ID
	if myUserID != kListDto.UserID {
		return nil, errors.New("don't have permission")
	}
	likeListDao, err := l.Dao.GetLikeListByUserId(ctx, kListDto)
	if err != nil {
		return nil, err
	}

	if len(likeListDao) == 0 {
		return nil, nil
	}

	likeVideoIdList := make([]uint, 0, len(likeListDao))
	for _, like := range likeListDao {
		likeVideoIdList = append(likeVideoIdList, like.VideoID)
	}
	likeVideoListDao, err := l.VideoDao.GetVideoListByVideoId(ctx, likeVideoIdList)
	if err != nil {
		return nil, err
	}
	likeVideoUserIdMap := make(map[uint]*models.User, len(likeVideoIdList))
	for _, like := range likeVideoListDao {
		likeVideoUserIdMap[like.AuthorID] = nil
	}
	likeVideoUserIdList := make([]uint, 0, len(likeVideoUserIdMap))
	for userId := range likeVideoUserIdMap {
		likeVideoUserIdList = append(likeVideoUserIdList, userId)
	}
	likeUserList, err := l.UserDao.GetUserListByIds(ctx, likeVideoUserIdList)
	if err != nil {
		return nil, err
	}
	for _, user := range likeUserList {
		likeVideoUserIdMap[user.ID] = user
	}
	var likeVideoList []*dto.Video
	for _, video := range likeVideoListDao {
		var likeVideo = new(dto.Video)
		_ = copier.Copy(likeVideo, video)
		var user = new(dto.User)
		_ = copier.Copy(user, likeVideoUserIdMap[video.AuthorID])
		likeVideo.Author = user
		likeVideoList = append(likeVideoList, likeVideo)
	}
	return likeVideoList, nil
}

// LikeAction 喜爱操作
func (l *LikeService) LikeAction(ctx context.Context, likeActionDto *dto.LikeActionDTO) error {
	if !l.VideoDao.IsExist(ctx, likeActionDto.VideoID) {
		return errors.New("video not exist")
	}

	switch likeActionDto.ActionType {
	case 1:
		return l.Dao.AddLike(ctx, likeActionDto)
	case 2:
		return l.Dao.CancelLike(ctx, likeActionDto)
	default:
		return errors.New("action type error")
	}
}
