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
func (m *MessageDao) AddMessage(ctx context.Context, dto *dto.AddMessageDTO) (*models.Message, error) {
	message := &models.Message{
		FromUserID: ctx.Value(global.LoginUser).(models.LoginUser).ID,
		ToUserID:   dto.ToUserID,
		Content:    dto.Content,
	}
	if err := m.DB.Model(&models.Message{}).WithContext(ctx).Create(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

// GetMessages 获取消息列表
func (m *MessageDao) GetMessages(
	ctx context.Context, dto *dto.MessageListDTO) (messages []*models.Message, preTime string, err error) {
	userID := ctx.Value(global.LoginUser).(models.LoginUser).ID
	if dto.PreMsgTime == "0" {
		preMsgTime := time.Now()
		err = m.DB.Model(&models.Message{}).WithContext(ctx).
			Where("(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
				dto.ToUserID, userID, userID, dto.ToUserID).Where("created_at <= ?", preMsgTime).
			Order("id").Find(&messages).Error
		if err != nil {
			return nil, "", err
		}
	} else {
		preMessageTime, err := strconv.ParseInt(dto.PreMsgTime, 10, 64)
		if err != nil {
			return nil, "", err
		}
		preMsgTime := time.UnixMilli(preMessageTime)
		err = m.DB.Model(&models.Message{}).WithContext(ctx).
			Where("(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
				dto.ToUserID, userID, userID, dto.ToUserID).
			Where("created_at > ?", preMsgTime).Order("id").Find(&messages).Error
		if err != nil {
			return nil, "", err
		}
	}
	if len(messages) == 0 {
		return nil, "0", nil
	}
	preTime = strconv.Itoa(int(messages[len(messages)-1].CreatedAt.UnixMilli()))
	return messages, preTime, nil
}

// GetUserOpenMsgList 获取用户开放联系人列表
func (m *MessageDao) GetUserOpenMsgList(ctx context.Context, dto *dto.OpenMsgListDTO) ([]*models.MessageOpen, error) {
	// 先查看是否有自己打开的联系人
	var openUsers []*models.MessageOpen
	err := m.DB.Model(&models.MessageOpen{}).WithContext(ctx).
		Where("user_id = ?", dto.UserId).Order("id DESC").Find(&openUsers).Error
	if err != nil {
		return nil, err
	}
	// 如果没有自己打开的联系人，则查看是否有其他用户传输消息
	//var otherOpenUsers []*models.MessageOpen
	//err = m.DB.Model(&models.MessageOpen{}).WithContext(ctx).
	//	Where("open_user_id = ?", dto.UserId).Order("id DESC").Find(&otherOpenUsers).Error
	//if err != nil {
	//	return nil, err
	//}
	//if len(openUsers) == 0 && len(otherOpenUsers) == 0 {
	//	return nil, nil
	//}
	//if len(otherOpenUsers) == 0 {
	//	return openUsers, nil
	//}
	//// 如果有其他用户传输消息
	//for _, otherUser := range otherOpenUsers {
	//	err := m.DB.Model(&models.Message{}).WithContext(ctx).
	//		Where("from_user_id = ? AND to_user_id = ?", otherUser.UserID, dto.UserId).
	//		First(&models.Message{}).Error
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		continue
	//	}
	//	if err != nil {
	//		return nil, err
	//	}
	//	openUsers = append(openUsers, &models.MessageOpen{
	//		UserID:     otherUser.OpenUserID,
	//		OpenUserID: otherUser.UserID,
	//	})
	//}
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
	return m.DB.WithContext(ctx).
		Where("user_id = ? AND open_user_id = ?", ctx.Value(global.LoginUser).(models.LoginUser).ID, openUserId).
		Delete(&models.MessageOpen{}).Error
}
