package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"

	"github.com/jinzhu/copier"

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
func (m *MessageService) SendMessage(ctx context.Context, addMsgDTO *dto.AddMessageDTO) (*dto.Message, error) {
	if !m.UserDao.IsExist(ctx, addMsgDTO.ToUserID) {
		return nil, errors.New("user not exist")
	}
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	// 添加消息
	messageDao, err := m.Dao.AddMessage(ctx, addMsgDTO)
	if err != nil {
		return nil, err
	}
	// 判断界面使用人数
	count, err := m.Dao.LinkUseCount(ctx, userId, addMsgDTO.ToUserID)
	if err != nil {
		return nil, err
	}
	if count == 1 {
		// 添加开放联系人
		err = m.Dao.AddOpenUser(ctx, addMsgDTO.ToUserID, userId, 2)
		if err != nil {
			return nil, err
		}
		// 更新使用人数
		err = m.Dao.UpdateLinkUseCount(ctx, userId, addMsgDTO.ToUserID, 2)
		if err != nil {
			return nil, err
		}
	}

	message := new(dto.Message)
	_ = copier.Copy(&message, &messageDao)

	fromUserDao, err := m.UserDao.GetUserById(ctx, ctx.Value(global.LoginUser).(models.LoginUser).ID)
	if err != nil {
		return nil, err
	}
	fromUser := new(dto.User)
	_ = copier.Copy(&fromUser, &fromUserDao)
	toUserDao, err := m.UserDao.GetUserById(ctx, addMsgDTO.ToUserID)
	if err != nil {
		return nil, err
	}
	toUser := new(dto.User)
	_ = copier.Copy(&toUser, &toUserDao)
	message.FromUser = fromUser
	message.ToUser = toUser
	return message, nil
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

	var formUserIdsMap = make(map[uint]*dto.User)
	var toUserIdsMap = make(map[uint]*dto.User)
	for _, msg := range msgDao {
		formUserIdsMap[msg.FromUserID] = nil
		toUserIdsMap[msg.ToUserID] = nil
	}
	formUserIds := make([]uint, 0, len(formUserIdsMap))
	for id := range formUserIdsMap {
		formUserIds = append(formUserIds, id)
	}
	toUserIds := make([]uint, 0, len(toUserIdsMap))
	for id := range toUserIdsMap {
		toUserIds = append(toUserIds, id)
	}
	formUsersDao, err := m.UserDao.GetUserListByIds(ctx, formUserIds)
	if err != nil {
		return nil, "", err
	}
	toUsersDao, err := m.UserDao.GetUserListByIds(ctx, toUserIds)
	if err != nil {
		return nil, "", err
	}
	fromUsers := make([]*dto.User, 0, len(formUsersDao))
	_ = copier.Copy(&fromUsers, &formUsersDao)
	toUsers := make([]*dto.User, 0, len(toUsersDao))
	_ = copier.Copy(&toUsers, &toUsersDao)
	for _, fromUser := range fromUsers {
		formUserIdsMap[fromUser.ID] = fromUser
	}
	for _, toUser := range toUsers {
		toUserIdsMap[toUser.ID] = toUser
	}

	var msgList []*dto.Message
	for _, msg := range msgDao {
		msgList = append(msgList, &dto.Message{
			FromUser:  formUserIdsMap[msg.FromUserID],
			ToUser:    toUserIdsMap[msg.ToUserID],
			Content:   msg.Content,
			CreatedAt: msg.CreatedAt.Unix(),
		})
	}
	return msgList, preMsgTime, nil
}

// GetUserOpenMsgList 获取开放联系人列表
func (m *MessageService) GetUserOpenMsgList(
	ctx context.Context, openMsgListDTO *dto.OpenMsgListDTO,
) ([]*dto.User, error) {
	msgOpenUsers, err := m.Dao.GetUserOpenMsgList(ctx, openMsgListDTO)
	if err != nil {
		return nil, err
	}
	if len(msgOpenUsers) == 0 {
		return nil, nil
	}
	var userIds []uint
	for _, user := range msgOpenUsers {
		userIds = append(userIds, user.OpenUserID)
	}
	usersDao, err := m.UserDao.GetUserListByIds(ctx, userIds)
	if err != nil {
		return nil, err
	}
	users := make([]*dto.User, 0, len(usersDao))
	_ = copier.Copy(&users, &usersDao)
	return users, nil
}

// AddOpenUser 添加开放联系人
func (m *MessageService) AddOpenUser(ctx context.Context, addDto *dto.AddOpenUserDTO) (*dto.User, error) {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	// 查询对方的使用此聊天界面的用户数
	count, err := m.Dao.LinkUseCount(ctx, addDto.OpenUserID, userId)
	if err != nil {
		return nil, err
	}
	fmt.Println(count)
	var useCount uint
	switch count {
	case 2: // 如果双方都在使用，则不做任何操作
		return nil, nil
	case 1: // 如果双方有一个在使用，则更新另一个的使用人数
		err = m.Dao.UpdateLinkUseCount(ctx, addDto.OpenUserID, userId, 2)
		if err != nil {
			return nil, err
		}
		useCount = 2
	default:
		useCount = 1
	}
	err = m.Dao.AddOpenUser(ctx, userId, addDto.OpenUserID, useCount)
	if err != nil {
		return nil, err
	}
	userDao, err := m.UserDao.GetUserById(ctx, addDto.OpenUserID)
	if err != nil {
		return nil, err
	}
	user := new(dto.User)
	_ = copier.Copy(&user, &userDao)
	return user, nil
}

// DeleteOpenUser 删除开放联系人
func (m *MessageService) DeleteOpenUser(ctx context.Context, dto *dto.DeleteOpenUserDTO) error {
	userId := ctx.Value(global.LoginUser).(models.LoginUser).ID
	// 查询双方用户是否都关闭了消息
	count, err := m.Dao.LinkUseCount(ctx, userId, dto.OpenUserID)
	if err != nil {
		return err
	}
	switch count {
	case 1: // 如果对方已经关闭了消息，则删除消息
		err = m.Dao.DeleteMsg(ctx, userId, dto.OpenUserID)
		if err != nil {
			return err
		}
	case 2: // 如果双方有一个在使用，则更新另一个的使用人数
		err = m.Dao.UpdateLinkUseCount(ctx, dto.OpenUserID, userId, 1)
		if err != nil {
			return err
		}
	}
	return m.Dao.DeleteLinkUser(ctx, dto.OpenUserID)
}
