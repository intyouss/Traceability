package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitCollectRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup, adGroup *gin.RouterGroup) {
		collectApi := api.NewCollectApi()
		collectAuthGroup := auGroup.Group("collect")
		{
			// 获取用户喜爱视频列表
			collectAuthGroup.GET("/list", collectApi.GetCollectList)
			// 用户喜爱操作
			collectAuthGroup.POST("/action", collectApi.CollectAction)
		}
	})
}
