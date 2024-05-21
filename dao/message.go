package dao

import (
	"context"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var MessageDaoIns *MessageDao

type MessageDao struct {
	*BaseDao
}

func NewMessageDao() *MessageDao {
	if MessageDaoIns == nil {
		MessageDaoIns = &MessageDao{
			BaseDao: NewBaseDao(),
		}
	}
	return MessageDaoIns
}

// AddMessage 添加消息
func (m *MessageDao) AddMessage(ctx context.Context, msgDto *dto.AddMessageDTO) (*models.Message, error) {
	message := &models.Message{
		FromUserID: ctx.Value(global.LoginUser).(models.LoginUser).ID,
		ToUserID:   msgDto.ToUserID,
		Content:    msgDto.Content,
	}
	if err := m.DB.Model(&models.Message{}).WithContext(ctx).Create(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

// GetMessages 获取消息列表
func (m *MessageDao) GetMessages(
	ctx context.Context, msgDto *dto.MessageListDTO) (messages []*models.Message, preTime string, err error) {
	userID := ctx.Value(global.LoginUser).(models.LoginUser).ID
	delTime, err := m.GetDeleteTime(ctx, msgDto.ToUserID, userID)
	if err != nil {
		return nil, "", err
	}
	if msgDto.PreMsgTime == "0" {
		preMsgTime := time.Now()
		if delTime != 0 {
			err = m.DB.Model(&models.Message{}).WithContext(ctx).
				Where("(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
					msgDto.ToUserID, userID, userID, msgDto.ToUserID).
				Where("created_at <= ? AND created_at > ?", preMsgTime, time.UnixMilli(delTime)).
				Order("id").Find(&messages).Error
			if err != nil {
				return nil, "", err
			}
		} else {
			err = m.DB.Model(&models.Message{}).WithContext(ctx).
				Where("(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
					msgDto.ToUserID, userID, userID, msgDto.ToUserID).Where("created_at <= ?", preMsgTime).
				Order("id").Find(&messages).Error
			if err != nil {
				return nil, "", err
			}
		}
		if len(messages) == 0 {
			return nil, "0", nil
		}
	} else {
		preMessageTime, err := strconv.ParseInt(msgDto.PreMsgTime, 10, 64)
		if err != nil {
			return nil, "", err
		}
		preMsgTime := time.UnixMilli(preMessageTime)
		err = m.DB.Model(&models.Message{}).WithContext(ctx).
			Where("(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
				msgDto.ToUserID, userID, userID, msgDto.ToUserID).
			Where("created_at > ?", preMsgTime).Order("id").Find(&messages).Error
		if err != nil {
			return nil, "", err
		}
		if len(messages) == 0 {
			return nil, msgDto.PreMsgTime, nil
		}
	}
	preTime = strconv.Itoa(int(messages[len(messages)-1].CreatedAt.UnixMilli()))
	return messages, preTime, nil
}

// GetUserOpenMsgList 获取用户打开联系人列表
func (m *MessageDao) GetUserOpenMsgList(
	ctx context.Context, userId uint,
) ([]*models.MessageOpen, error) {
	var openUsers []*models.MessageOpen
	err := m.DB.Model(&models.MessageOpen{}).WithContext(ctx).
		Where("user_id = ?", userId).Order("id DESC").Find(&openUsers).Error
	if err != nil {
		return nil, err
	}
	return openUsers, nil
}

// LinkUseCount 双方聊天使用人数
func (m *MessageDao) LinkUseCount(ctx context.Context, userId, openUserId uint) (uint, error) {
	var openMsg models.MessageOpen
	err := m.DB.Model(&models.MessageOpen{}).WithContext(ctx).
		Where("user_id = ? AND open_user_id = ?", userId, openUserId).
		First(&openMsg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return openMsg.UseCount, nil
}

// UpdateLinkUseCount 更新双方聊天使用人数
func (m *MessageDao) UpdateLinkUseCount(ctx context.Context, userId, openUserId uint, count uint) error {
	return m.DB.Model(&models.MessageOpen{}).WithContext(ctx).
		Where("user_id = ? AND open_user_id = ?", userId, openUserId).
		Update("use_count", count).Error
}

// UpdateDeleteTime 更新删除时间
func (m *MessageDao) UpdateDeleteTime(ctx context.Context, userId, openUserId uint, delTime int64) error {
	return m.DB.Model(&models.MessageOpen{}).WithContext(ctx).
		Where("user_id = ? AND open_user_id = ?", userId, openUserId).
		Update("delete_time", delTime).Error
}

// GetDeleteTime 获取删除时间
func (m *MessageDao) GetDeleteTime(ctx context.Context, userId, openUserId uint) (int64, error) {
	var openMsg models.MessageOpen
	err := m.DB.Model(&models.MessageOpen{}).WithContext(ctx).
		Where("user_id = ? AND open_user_id = ?", userId, openUserId).
		First(&openMsg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return openMsg.DeleteTime, nil
}

// AddOpenUser 添加开放联系人
func (m *MessageDao) AddOpenUser(ctx context.Context, userId, openUserId uint, count uint) error {
	openMsg := &models.MessageOpen{
		UserID:     userId,
		OpenUserID: openUserId,
		UseCount:   count,
	}
	return m.DB.Model(&models.MessageOpen{}).WithContext(ctx).Create(&openMsg).Error
}

// DeleteMsg 删除消息
func (m *MessageDao) DeleteMsg(ctx context.Context, userId, openUserId uint) error {
	return m.DB.WithContext(ctx).
		Where("from_user_id = ? AND to_user_id = ?", userId, openUserId).
		Or("from_user_id = ? AND to_user_id = ?", openUserId, userId).
		Delete(&models.Message{}).Error
}

// DeleteLinkUser 删除开放联系人
func (m *MessageDao) DeleteLinkUser(ctx context.Context, openUserId uint) error {
	return m.DB.WithContext(ctx).Unscoped().
		Where("user_id = ? AND open_user_id = ?", ctx.Value(global.LoginUser).(models.LoginUser).ID, openUserId).
		Delete(&models.MessageOpen{}).Error
}
