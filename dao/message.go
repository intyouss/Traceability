package dao

import (
	"context"
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
func (m *MessageDao) AddMessage(ctx context.Context, dto *dto.AddMessageDTO) error {
	message := &models.Message{
		FromUserID: ctx.Value(global.LoginUser).(uint),
		ToUserID:   dto.ToUserID,
		Content:    dto.Content,
	}
	return m.DB.Model(&models.Message{}).WithContext(ctx).Create(&message).Error
}

// GetMessages 获取消息列表
func (m *MessageDao) GetMessages(ctx context.Context, dto *dto.MessageListDTO) ([]*models.Message, error) {
	var messages []*models.Message
	err := m.DB.Model(&models.Message{}).WithContext(ctx).
		Where("to_user_id = ?", dto.ToUserID).Find(&messages).Error
	return messages, err
}
