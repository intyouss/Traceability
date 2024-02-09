package dao

import (
	"context"
	"time"

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
		FromUserID: ctx.Value(global.LoginUser).(models.LoginUser).ID,
		ToUserID:   dto.ToUserID,
		Content:    dto.Content,
	}
	return m.DB.Model(&models.Message{}).WithContext(ctx).Create(&message).Error
}

// GetMessages 获取消息列表
func (m *MessageDao) GetMessages(ctx context.Context, dto *dto.MessageListDTO) (messages []*models.Message, err error) {
	userID := ctx.Value(global.LoginUser).(models.LoginUser).ID
	if dto.PreMsgTime != nil && *dto.PreMsgTime == 0 {
		err = m.DB.Model(&models.Message{}).WithContext(ctx).
			Where("(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
				dto.ToUserID, userID, userID, dto.ToUserID).
			Order("id DESC").Find(&messages).Error
	} else {
		err = m.DB.Model(&models.Message{}).WithContext(ctx).
			Where("(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
				dto.ToUserID, userID, userID, dto.ToUserID).
			Where("created_at >= ?", time.Unix(*dto.PreMsgTime, 0)).
			Order("id DESC").Find(&messages).Error
	}
	return
}
