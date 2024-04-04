package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/utils"

	"github.com/jinzhu/copier"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var VideoServiceIns *VideoService

type VideoService struct {
	BaseService
	UserDao     *dao.UserDao
	LikeDao     *dao.LikeDao
	CollectDao  *dao.CollectDao
	RelationDao *dao.RelationDao
	Dao         *dao.VideoDao
}

func NewVideoService() *VideoService {
	if VideoServiceIns == nil {
		VideoServiceIns = &VideoService{
			Dao:         dao.NewVideoDao(),
			UserDao:     dao.NewUserDao(),
			LikeDao:     dao.NewLikeDao(),
			CollectDao:  dao.NewCollectDao(),
			RelationDao: dao.NewRelationDao(),
		}
	}
	return VideoServiceIns
}

// GetVideoList 获取视频列表
func (v *VideoService) GetVideoList(
	ctx context.Context, vListDTO *dto.VideoListDTO,
) (videos []*dto.Video, nextTime string, err error) {
	var videosDao []*models.Video
	// 分类获取视频列表
	switch vListDTO.Type {
	case 1:
		videosDao, nextTime, err = v.Dao.GetVideoList(ctx, vListDTO)
	case 2:
		videosDao, nextTime, err = v.Dao.GetFocusVideoList(ctx, vListDTO)
	case 3:
		videosDao, nextTime, err = v.Dao.GetFriendVideoList(ctx, vListDTO)
	case 4:
		videosDao, nextTime, err = v.Dao.GetVideoList(ctx, vListDTO)
	default:
		return nil, "", errors.New("type error")
	}
	if err != nil {
		return nil, "", err
	}
	if len(videosDao) == 0 {
		return nil, "0", nil
	}
	for _, video := range videosDao {
		err = v.UpdateUrls(ctx, video)
		if err != nil {
			return nil, "", err
		}
	}

	// 获取作者信息
	authorMap := make(map[uint]*models.User)
	for _, video := range videosDao {
		authorMap[video.AuthorID] = nil
	}

	var authorIds []uint
	for authorId := range authorMap {
		authorIds = append(authorIds, authorId)
	}

	authors, err := v.UserDao.GetUserListByIds(ctx, authorIds)
	if err != nil {
		return nil, "", err
	}
	for _, author := range authors {
		authorMap[author.ID] = author
	}

	// 视频判断是否已点赞 是否已收藏， 作者判断是否已关注
	likeMap := make(map[uint]bool)
	collectMap := make(map[uint]bool)
	focusMap := make(map[uint]bool)
	var wg sync.WaitGroup
	errChan := make(chan error)
	if ctx.Value(global.LoginUser) != nil {
		videoIds := make([]uint, 0, len(videosDao))
		for _, video := range videosDao {
			videoIds = append(videoIds, video.ID)
		}
		wg.Add(1)
		go func(vi []uint) {
			defer wg.Done()
			likeMap, err = v.LikeDao.IsLikedByList(ctx, vi)
			if err != nil {
				errChan <- err
			}
		}(videoIds)

		wg.Add(1)
		go func(vi []uint) {
			defer wg.Done()
			collectMap, err = v.CollectDao.IsCollectedByList(ctx, vi)
			if err != nil {
				errChan <- err
			}
		}(videoIds)

		wg.Add(1)
		go func(au []uint) {
			defer wg.Done()
			focusMap, err = v.RelationDao.IsFocusedByList(ctx, au)
			if err != nil {
				errChan <- err
			}
		}(authorIds)
		wg.Wait()
		select {
		case err = <-errChan:
			return nil, "", err
		default:
		}
	}

	// 组装数据
	for _, video := range videosDao {
		var videoDTO = new(dto.Video)
		_ = copier.Copy(videoDTO, video)
		videoDTO.CreatedAt = utils.TimeFormat(video.CreatedAt)
		var user = new(dto.User)
		_ = copier.Copy(user, authorMap[video.AuthorID])
		videoDTO.Author = user
		if ctx.Value(global.LoginUser) != nil {
			videoDTO.IsLike = likeMap[video.ID]
			videoDTO.IsCollect = collectMap[video.ID]
			videoDTO.Author.IsFocus = focusMap[video.AuthorID]
		} else {
			videoDTO.IsLike = false
			videoDTO.IsCollect = false
			videoDTO.Author.IsFocus = false
		}
		videos = append(videos, videoDTO)
	}
	return videos, nextTime, nil
}

// GetVideoListByUserId 根据用户id获取视频列表
func (v *VideoService) GetVideoListByUserId(
	ctx context.Context, uVideoListDto *dto.UserVideoListDTO,
) ([]*dto.Video, error) {
	if !v.UserDao.IsExist(ctx, uVideoListDto.UserID) && ctx.
		Value(global.LoginUser).(models.LoginUser).ID != uVideoListDto.UserID {
		return nil, errors.New("user not exist")
	}
	// 获取视频列表
	videosDao, err := v.Dao.GetVideoListByUserId(ctx, uVideoListDto)
	if err != nil {
		return nil, err
	}
	if len(videosDao) == 0 {
		return nil, nil
	}
	for _, video := range videosDao {
		err = v.UpdateUrls(ctx, video)
		if err != nil {
			return nil, err
		}
	}

	// 获取作者信息
	var userDao *models.User
	userDao, err = v.UserDao.GetUserById(ctx, uVideoListDto.UserID)
	if err != nil {
		return nil, err
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)

	// 视频判断是否已点赞 是否已收藏， 作者判断是否已关注
	likeMap := make(map[uint]bool)
	collectMap := make(map[uint]bool)
	var isFocus bool
	var wg sync.WaitGroup
	errChan := make(chan error)
	if ctx.Value(global.LoginUser) != nil {
		videoIds := make([]uint, 0, len(videosDao))
		for _, video := range videosDao {
			videoIds = append(videoIds, video.ID)
		}
		wg.Add(1)
		go func(vi []uint) {
			defer wg.Done()
			likeMap, err = v.LikeDao.IsLikedByList(ctx, vi)
			if err != nil {
				errChan <- err
			}
		}(videoIds)

		wg.Add(1)
		go func(vi []uint) {
			defer wg.Done()
			collectMap, err = v.CollectDao.IsCollectedByList(ctx, vi)
			if err != nil {
				errChan <- err
			}
		}(videoIds)

		wg.Add(1)
		go func() {
			defer wg.Done()
			isFocus, err = v.RelationDao.IsFocused(ctx, uVideoListDto.UserID)
			if err != nil {
				errChan <- err
			}
		}()
		wg.Wait()
		select {
		case err = <-errChan:
			return nil, err
		default:
		}
	}

	var videos = make([]*dto.Video, 0, len(videosDao))
	for _, video := range videosDao {
		var videoDTO = new(dto.Video)
		_ = copier.Copy(videoDTO, video)
		videoDTO.CreatedAt = utils.TimeFormat(video.CreatedAt)
		videoDTO.Author = user
		if ctx.Value(global.LoginUser) != nil {
			videoDTO.IsLike = likeMap[video.ID]
			videoDTO.IsCollect = collectMap[video.ID]
			videoDTO.Author.IsFocus = isFocus
		} else {
			videoDTO.IsLike = false
			videoDTO.IsCollect = false
			videoDTO.Author.IsFocus = false
		}
		videos = append(videos, videoDTO)
	}
	return videos, nil
}

// GetVideoSearch 获取视频搜索结果
func (v *VideoService) GetVideoSearch(
	ctx context.Context, searchDTO *dto.VideoSearchDTO,
) (videos []*dto.Video, err error) {
	var videosDao []*models.Video
	// 分类获取视频列表
	switch searchDTO.Type {
	case 1:
		videosDao, err = v.Dao.GetVideoSearchByAuthorAndTitle(ctx, searchDTO)
	case 2:
		videosDao, err = v.Dao.GetVideoSearchByTitle(ctx, searchDTO)
	}
	if err != nil {
		return nil, err
	}
	if len(videosDao) == 0 {
		return nil, nil
	}
	for _, video := range videosDao {
		err = v.UpdateUrls(ctx, video)
		if err != nil {
			return nil, err
		}
	}

	// 获取作者信息
	var authorIds = make([]uint, 0, len(videosDao))
	for _, video := range videosDao {
		authorIds = append(authorIds, video.AuthorID)
	}
	var authorIdsMap = make(map[uint]*models.User)
	for _, video := range videosDao {
		authorIdsMap[video.AuthorID] = nil
	}
	authors, err := v.UserDao.GetUserListByIds(ctx, authorIds)
	if err != nil {
		return nil, err
	}
	for _, author := range authors {
		authorIdsMap[author.ID] = author
	}

	// 视频判断是否已点赞 是否已收藏， 作者判断是否已关注
	likeMap := make(map[uint]bool)
	collectMap := make(map[uint]bool)
	focusMap := make(map[uint]bool)
	var wg sync.WaitGroup
	errChan := make(chan error)
	if ctx.Value(global.LoginUser) != nil {
		videoIds := make([]uint, 0, len(videosDao))
		for _, video := range videosDao {
			videoIds = append(videoIds, video.ID)
		}
		wg.Add(1)
		go func(vi []uint) {
			defer wg.Done()
			likeMap, err = v.LikeDao.IsLikedByList(ctx, vi)
			if err != nil {
				errChan <- err
			}
		}(videoIds)

		wg.Add(1)
		go func(vi []uint) {
			defer wg.Done()
			collectMap, err = v.CollectDao.IsCollectedByList(ctx, vi)
			if err != nil {
				errChan <- err
			}
		}(videoIds)

		wg.Add(1)
		go func(au []uint) {
			defer wg.Done()
			focusMap, err = v.RelationDao.IsFocusedByList(ctx, au)
			if err != nil {
				errChan <- err
			}
		}(authorIds)
		wg.Wait()
		select {
		case err = <-errChan:
			return nil, err
		default:
		}
	}

	for _, video := range videosDao {
		var videoDTO = new(dto.Video)
		_ = copier.Copy(videoDTO, video)
		videoDTO.CreatedAt = utils.TimeFormat(video.CreatedAt)
		var user = new(dto.User)
		_ = copier.Copy(user, authorIdsMap[video.AuthorID])
		videoDTO.Author = user
		if ctx.Value(global.LoginUser) != nil {
			videoDTO.IsLike = likeMap[video.ID]
			videoDTO.IsCollect = collectMap[video.ID]
			videoDTO.Author.IsFocus = focusMap[video.AuthorID]
		} else {
			videoDTO.IsLike = false
			videoDTO.IsCollect = false
			videoDTO.Author.IsFocus = false
		}
		videos = append(videos, videoDTO)
	}
	return videos, nil
}

// UpdateUrls 更新url
func (v *VideoService) UpdateUrls(ctx context.Context, video *models.Video) error {
	firstOk, err := v.Dao.CheckUrl(video.PlayUrl)
	if err != nil {
		return err
	}
	secondOk, err := v.Dao.CheckUrl(video.CoverUrl)
	if err != nil {
		return err
	}
	if firstOk {
		playUrl, err := v.Dao.GetRemoteVideoUrl(ctx, video.Title)
		if err != nil {
			return err
		}
		video.PlayUrl = playUrl
	}
	if secondOk {
		coverUrl, err := v.Dao.GetRemoteCoverImageUrl(ctx, video.Title)
		if err != nil {
			return err
		}
		video.CoverUrl = coverUrl
	}
	if firstOk || secondOk {
		go func(vi *models.Video) {
			err = v.Dao.UpdateDBUrl(ctx, vi.ID, vi.PlayUrl, vi.CoverUrl)
			if err != nil {
				v.logger.Error(err)
			}
		}(video)
	}
	return nil
}

func (v *VideoService) GetVideoInfo(ctx context.Context, idDto *dto.CommonIDDTO) (*dto.Video, error) {
	video, err := v.Dao.GetVideoByVideoId(ctx, *idDto.ID)
	if err != nil {
		return nil, err
	}
	var videoDTO = new(dto.Video)
	_ = copier.Copy(videoDTO, video)
	author, err := v.UserDao.GetUserById(ctx, video.AuthorID)
	if err != nil {
		return nil, err
	}
	var user = new(dto.User)
	_ = copier.Copy(user, author)
	videoDTO.Author = user
	videoDTO.CreatedAt = utils.TimeFormat(video.CreatedAt)
	if ctx.Value(global.LoginUser) != nil {
		like, err := v.LikeDao.IsLiked(ctx, *idDto.ID)
		if err != nil {
			return nil, err
		}
		collect, err := v.CollectDao.IsCollected(ctx, *idDto.ID)
		if err != nil {
			return nil, err
		}
		focus, err := v.RelationDao.IsFocused(ctx, video.AuthorID)
		if err != nil {
			return nil, err
		}
		fmt.Println(like, collect, focus)
		videoDTO.IsLike = like
		videoDTO.IsCollect = collect
		videoDTO.Author.IsFocus = focus
	} else {
		videoDTO.IsLike = false
		videoDTO.IsCollect = false
		videoDTO.Author.IsFocus = false
	}
	return videoDTO, nil
}

// UploadVideo 上传视频
func (v *VideoService) UploadVideo(ctx context.Context, upload *dto.VideoUploadDTO) (string, error) {
	err := v.Dao.UploadVideo(ctx, upload)
	if err != nil {
		return "", err
	}
	// 获取视频和封面的url
	return v.Dao.GetRemoteVideoUrl(ctx, upload.Title)
}

// UploadImage 上传视频封面
func (v *VideoService) UploadImage(ctx context.Context, upload *dto.ImageUploadDTO) (string, error) {
	err := v.Dao.UploadCoverImage(ctx, upload)
	if err != nil {
		return "", err
	}
	return v.Dao.GetRemoteCoverImageUrl(ctx, upload.Title)
}

// SaveVideoInfo 保存视频
func (v *VideoService) SaveVideoInfo(ctx context.Context, publishDTO *dto.PublishDTO) error {
	err := v.Dao.SaveVideoInfo(ctx, publishDTO)
	if err != nil {
		return err
	}
	year := uint(time.Now().Year())
	month := uint(time.Now().Month())
	day := uint(time.Now().Day())

	// 更新视频发布日增长数
	ok, _, err := v.Dao.GetVideoIncrease(ctx, year, month, day)
	if err != nil {
		return err
	}
	if ok {
		err = v.Dao.UpdateVideoIncreaseCount(ctx, year, month, day, 1)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteRemoteVideo 删除远程视频信息
func (v *VideoService) DeleteRemoteVideo(ctx context.Context, abolishDTO *dto.AbolishVideoUploadDTO) error {
	var deleteImg bool
	var deleteVideo bool
	switch abolishDTO.Type {
	case 1:
		deleteImg = true
		deleteVideo = true
	case 2:
		deleteImg = false
		deleteVideo = true
	case 3:
		deleteImg = true
		deleteVideo = false
	}
	var wg sync.WaitGroup
	errChan := make(chan error)
	// 删除封面
	wg.Add(1)
	go func() {
		defer wg.Done()
		if deleteImg {
			err := v.Dao.DeleteRemoteCoverImage(ctx, abolishDTO.Title)
			if err != nil {
				errChan <- err
				return
			}
		}
	}()
	// 删除视频
	wg.Add(1)
	go func() {
		defer wg.Done()
		if deleteVideo {
			err := v.Dao.DeleteRemoteVideo(ctx, abolishDTO.Title)
			if err != nil {
				errChan <- err
				return
			}
		}
	}()
	wg.Wait()
	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}
}

// GetVideoIncrease 获取月总日视频发布数增长列表
func (v *VideoService) GetVideoIncrease(
	ctx context.Context, timeDTO *dto.VideoIncreaseListDTO,
) ([]*dto.VideoIncrease, error) {
	list, err := v.Dao.GetVideoIncreaseList(ctx, timeDTO.Year, timeDTO.Month)
	if err != nil {
		return nil, err
	}
	var IncreaseList = make([]*dto.VideoIncrease, 0, len(list))
	_ = copier.Copy(&IncreaseList, &list)
	return IncreaseList, nil
}
