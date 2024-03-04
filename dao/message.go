package dao

import (
	"context"
	"strconv"
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
func (m *MessageDao) GetMessages(
	ctx context.Context, dto *dto.MessageListDTO) (messages []*models.Message, preTime string, err error) {
	userID := ctx.Value(global.LoginUser).(models.LoginUser).ID
	if dto.PreMsgTime == "0" {
		preMsgTime := time.Now()
		err = m.DB.Model(&models.Message{}).WithContext(ctx).
			Where("(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
				dto.ToUserID, userID, userID, dto.ToUserID).Where("created_at <= ?", preMsgTime).
			Order("id").Find(&messages).Error
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
	if len(messages) != 0 {
		preTime = strconv.Itoa(int(messages[len(messages)-1].CreatedAt.UnixMilli()))
	}
	return
}
