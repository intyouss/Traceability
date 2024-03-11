package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/service"
	"github.com/intyouss/Traceability/service/dto"
)

const (
	ErrCodeSendMessage = iota + 60001
	ErrCodeGetMessage
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

	if err := m.Service.SendMessage(ctx, &addMsgDTO); err != nil {
		m.Fail(&Response{Code: ErrCodeSendMessage, Msg: err.Error()})
		return
	}
	m.Success(&Response{})

}

// GetMessages 获取消息列表
// @Summary 获取消息列表
// @Description 获取消息列表
// @Param token header string true "token"
// @Param to_user_id formData int true "目标用户id"
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
				"pre_msg_time": 0,
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
