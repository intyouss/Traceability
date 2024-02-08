package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitMessageRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup) {
		messageApi := api.NewMessageApi()
		messageAuthGroup := auGroup.Group("message")
		{
			// 发送消息
			messageAuthGroup.POST("/send", messageApi.SendMessage)
			// 获取消息列表
			messageAuthGroup.GET("/chat", messageApi.GetMessages)
		}
	})
}
