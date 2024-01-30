package service

import (
	"context"
	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/service/dto"
)

var MessageServiceIns *MessageService

type MessageService struct {
	BaseService
	Dao *dao.MessageDao
}

func NewMessageService() *MessageService {
	if MessageServiceIns == nil {
		MessageServiceIns = &MessageService{
			Dao: dao.NewMessageDao(),
		}
	}
	return MessageServiceIns
}

// SendMessage 发送消息
func (m *MessageService) SendMessage(ctx context.Context, dto *dto.AddMessageDTO) error {
	return m.Dao.AddMessage(ctx, dto)
}

// GetMessages 获取消息列表
func (m *MessageService) GetMessages(ctx context.Context, dto *dto.MessageListDTO) ([]*models.Message, error) {
	return m.Dao.GetMessages(ctx, dto)
}
