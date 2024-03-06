package service

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/intyouss/Traceability/global"

	"github.com/jinzhu/copier"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var VideoServiceIns *VideoService

type VideoService struct {
	BaseService
	UserDao *dao.UserDao
	Dao     *dao.VideoDao
}

func NewVideoService() *VideoService {
	if VideoServiceIns == nil {
		VideoServiceIns = &VideoService{
			Dao:     dao.NewVideoDao(),
			UserDao: dao.NewUserDao(),
		}
	}
	return VideoServiceIns
}

// GetVideoList 获取视频列表
func (v *VideoService) GetVideoList(
	ctx context.Context, vListDTO *dto.VideoListDTO,
) (videos []*dto.Video, nextTime string, err error) {
	var videosDao []*models.Video
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
	err = v.Dao.UpdateUrl(ctx, videosDao)
	if err != nil {
		return nil, "", err
	}

	authorMap := make(map[uint]*models.User)
	for _, video := range videosDao {
		authorMap[video.AuthorID] = nil
	}
	var authorIds []uint
	for authorId := range authorMap {
		authorIds = append(authorIds, authorId)
	}

	// 调用userApi
	authors, err := v.UserDao.GetUserListByIds(ctx, authorIds)
	if err != nil {
		return nil, "", err
	}
	for _, author := range authors {
		authorMap[author.ID] = author
	}

	// 组装数据
	for _, video := range videosDao {
		var videoDTO = new(dto.Video)
		_ = copier.Copy(videoDTO, video)
		videoDTO.CreatedAt = timeFormat(video.CreatedAt)
		var user = new(dto.User)
		_ = copier.Copy(user, authorMap[video.AuthorID])
		videoDTO.Author = user
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
	videosDao, err := v.Dao.GetVideoListByUserId(ctx, uVideoListDto)
	if err != nil {
		return nil, err
	}
	if len(videosDao) == 0 {
		return nil, nil
	}
	err = v.Dao.UpdateUrl(ctx, videosDao)
	if err != nil {
		return nil, err
	}

	var userDao *models.User
	userDao, err = v.UserDao.GetUserById(ctx, uVideoListDto.UserID)
	if err != nil {
		return nil, err
	}
	var user = new(dto.User)
	_ = copier.Copy(user, userDao)

	var videos = make([]*dto.Video, 0, len(videosDao))
	for _, video := range videosDao {
		var videoDTO = new(dto.Video)
		_ = copier.Copy(videoDTO, video)
		videoDTO.CreatedAt = timeFormat(video.CreatedAt)
		videoDTO.Author = user
		videos = append(videos, videoDTO)
	}
	return videos, nil
}

// GetVideoSearch 获取视频搜索结果
func (v *VideoService) GetVideoSearch(
	ctx context.Context, searchDTO *dto.VideoSearchDTO,
) (videos []*dto.Video, err error) {
	var videosDao []*models.Video
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
	err = v.Dao.UpdateUrl(ctx, videosDao)
	if err != nil {
		return nil, err
	}

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
	for _, video := range videosDao {
		var videoDTO = new(dto.Video)
		_ = copier.Copy(videoDTO, video)
		videoDTO.CreatedAt = timeFormat(video.CreatedAt)
		var user = new(dto.User)
		_ = copier.Copy(user, authorIdsMap[video.AuthorID])
		videoDTO.Author = user
		videos = append(videos, videoDTO)
	}
	return videos, nil
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

// timeFormat 时间格式化
func timeFormat(t time.Time) string {
	since := time.Since(t)
	switch {
	case since < time.Minute: // 如果是一分钟内的时间，返回刚刚
		return "刚刚"
	case since < time.Hour: // 如果是一小时内的时间，返回是几分钟前
		return strings.Split(since.String(), "m")[0] + "分钟前"
	case since < 24*time.Hour: // 如果是超过一个小时的时间，返回是几小时前
		return strings.Split(since.String(), "h")[0] + "小时前"
	case since < 7*24*time.Hour: // 如果超过一天但是在一周内，返回是几天前
		x, _ := strconv.Atoi(strings.Split(since.String(), "h")[0])
		return strconv.Itoa(x/24) + "天前"
	case since < 21*24*time.Hour: // 如果超过一周，但不超过三周，返回是几周前
		x, _ := strconv.Atoi(strings.Split(since.String(), "h")[0])
		return strconv.Itoa(x/(7*24)) + "周前"
	default: // 如果超过三周，返回年月日
		return t.Format("2006-01-02")
	}
}
