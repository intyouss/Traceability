package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
)

const (
	ErrCodeSendMessage = iota + 60001
	ErrCodeGetMessage
	ErrCodeGetUserOpenMsgList
	ErrCodeAddOpenUser
	ErrCodeDeleteOpenUser
)

type MessageApi struct {
	BaseApi
	Service *service.MessageService
}

func NewMessageApi() MessageApi {
	return MessageApi{
		BaseApi: NewBaseApi(),
		Service: service.NewMessageService(),
	}
}

// SendMessage 发送消息
// @Summary 发送消息
// @Description 发送消息(暂时设置为不互相关注也能发消息)
// @Param token header string true "token"
// @Param to_user_id formData int true "目标用户id"
// @Param content formData string true "消息内容"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/message/send [post]
func (m MessageApi) SendMessage(ctx *gin.Context) {
	var addMsgDTO dto.AddMessageDTO
	err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &addMsgDTO}).GetError()
	if err != nil {
		m.Fail(&Response{Code: ErrCodeSendMessage, Msg: err.Error()})
		return
	}
	message, err := m.Service.SendMessage(ctx, &addMsgDTO)
	if err != nil {
		m.Fail(&Response{Code: ErrCodeSendMessage, Msg: err.Error()})
		return
	}

	m.Success(&Response{
		Data: gin.H{
			"message": message,
		},
	})

}

// GetMessages 获取消息列表
// @Summary 获取消息列表
// @Description 获取消息列表
// @Param token header string true "token"
// @Param to_user_id query int true "目标用户id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/message/chat [get]
func (m MessageApi) GetMessages(ctx *gin.Context) {
	var msgListDTO dto.MessageListDTO
	err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &msgListDTO}).GetError()
	if err != nil {
		m.Fail(&Response{Code: ErrCodeGetMessage, Msg: err.Error()})
		return
	}

	msgList, preMsgTime, err := m.Service.GetMessages(ctx, &msgListDTO)
	if err != nil {
		m.Fail(&Response{Code: ErrCodeGetMessage, Msg: err.Error()})
		return
	}
	if len(msgList) == 0 {
		m.Success(&Response{
			Data: gin.H{
				"messages":     []*dto.Message{},
				"pre_msg_time": preMsgTime,
			},
		})
		return
	}

	m.Success(&Response{
		Data: gin.H{
			"messages":     msgList,
			"pre_msg_time": preMsgTime,
		},
	})
}

// GetUserOpenMsgList 获取用户开放消息列表
// @Summary 获取用户开放消息列表
// @Description 获取用户开放消息列表
// @Param token header string true "token"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/message/open [get]
func (m MessageApi) GetUserOpenMsgList(ctx *gin.Context) {
	var openListDTO dto.OpenMsgListDTO
	err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &openListDTO}).GetError()
	if err != nil {
		m.Fail(&Response{Code: ErrCodeGetUserOpenMsgList, Msg: err.Error()})
		return
	}

	users, err := m.Service.GetUserOpenMsgList(ctx, &openListDTO)
	if err != nil {
		m.Fail(&Response{Code: ErrCodeGetUserOpenMsgList, Msg: err.Error()})
		return
	}
	if len(users) == 0 {
		m.Success(&Response{
			Data: gin.H{
				"users": []*dto.User{},
			},
		})
		return
	}

	m.Success(&Response{
		Data: gin.H{
			"users": users,
		},
	})
}

// AddOpenUser 添加开放联系人
// @Summary 添加开放联系人
// @Description 添加开放联系人
// @Param token header string true "token"
// @Param open_user_id formData int true "开放联系人id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/message/open/add [post]
func (m MessageApi) AddOpenUser(ctx *gin.Context) {
	var addOpenUserDTO dto.AddOpenUserDTO
	err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &addOpenUserDTO}).GetError()
	if err != nil {
		m.Fail(&Response{Code: ErrCodeAddOpenUser, Msg: err.Error()})
		return
	}
	user, err := m.Service.AddOpenUser(ctx, &addOpenUserDTO)
	if err != nil {
		m.Fail(&Response{Code: ErrCodeAddOpenUser, Msg: err.Error()})
		return
	}
	if user == nil {
		m.Success(&Response{})
		return
	}

	m.Success(&Response{
		Data: gin.H{
			"user": user,
		},
	})
}

// DeleteOpenUser 删除开放联系人
// @Summary 删除开放联系人
// @Description 删除开放联系人
// @Param token header string true "token"
// @Param open_user_id formData int true "开放联系人id"
// @Success 200 {string} Response
// @Failure 400 {string} Response
// @Router /api/v1/message/open/delete [post]
func (m MessageApi) DeleteOpenUser(ctx *gin.Context) {
	var delOpenUserDTO dto.DeleteOpenUserDTO
	err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &delOpenUserDTO}).GetError()
	if err != nil {
		m.Fail(&Response{Code: ErrCodeDeleteOpenUser, Msg: err.Error()})
		return
	}
	fmt.Println(delOpenUserDTO.OpenUserID)
	err = m.Service.DeleteOpenUser(ctx, &delOpenUserDTO)
	if err != nil {
		m.Fail(&Response{Code: ErrCodeDeleteOpenUser, Msg: err.Error()})
		return
	}

	m.Success(&Response{})
}
