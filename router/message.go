package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitMessageRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup, adGroup *gin.RouterGroup) {
		messageApi := api.NewMessageApi()
		messageAuthGroup := auGroup.Group("message")
		{
			// 发送消息
			messageAuthGroup.POST("/send", messageApi.SendMessage)
			// 获取消息列表
			messageAuthGroup.GET("/chat", messageApi.GetMessages)
			// 获取开放联系人列表
			messageAuthGroup.GET("/open", messageApi.GetUserOpenMsgList)
			// 添加开放联系人消息
			messageAuthGroup.POST("/open/add", messageApi.AddOpenUser)
			// 删除开放联系人消息
			messageAuthGroup.POST("/open/delete", messageApi.DeleteOpenUser)
		}
	})
}
