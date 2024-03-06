package service

import (
	"context"
	"errors"

	"github.com/intyouss/Traceability/dao"
	"github.com/intyouss/Traceability/service/dto"
)

var MessageServiceIns *MessageService

type MessageService struct {
	BaseService
	UserDao *dao.UserDao
	Dao     *dao.MessageDao
}

func NewMessageService() *MessageService {
	if MessageServiceIns == nil {
		MessageServiceIns = &MessageService{
			Dao:     dao.NewMessageDao(),
			UserDao: dao.NewUserDao(),
		}
	}
	return MessageServiceIns
}

// SendMessage 发送消息
func (m *MessageService) SendMessage(ctx context.Context, addMsgDTO *dto.AddMessageDTO) error {
	if !m.UserDao.IsExist(ctx, addMsgDTO.ToUserID) {
		return errors.New("user not exist")
	}
	return m.Dao.AddMessage(ctx, addMsgDTO)
}

// GetMessages 获取消息列表
func (m *MessageService) GetMessages(
	ctx context.Context, msgListDTO *dto.MessageListDTO,
) ([]*dto.Message, string, error) {
	if !m.UserDao.IsExist(ctx, msgListDTO.ToUserID) {
		return nil, "", errors.New("user not exist")
	}

	msgDao, preMsgTime, err := m.Dao.GetMessages(ctx, msgListDTO)
	if err != nil {
		return nil, "", err
	}
	if len(msgDao) == 0 {
		return nil, "", nil
	}
	var msgList []*dto.Message
	for _, msg := range msgDao {
		msgList = append(msgList, &dto.Message{
			FromUserID: msg.FromUserID,
			ToUserID:   msg.ToUserID,
			Content:    msg.Content,
			CreatedAt:  msg.CreatedAt.Unix(),
		})
	}
	return msgList, preMsgTime, nil
}
