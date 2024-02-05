package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitVideoRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup) {
		videoApi := api.NewVideoApi()
		videoDefaultGroup := dfGroup.Group("video")
		{
			// 视频feed流
			videoDefaultGroup.POST("/feed", videoApi.GetVideoFeed)
			// 用户发布视频列表
			videoDefaultGroup.GET("/list", videoApi.GetUserVideoList)
		}
		videoAuthGroup := auGroup.Group("video")
		{
			// 发布视频
			videoAuthGroup.POST("/publish", videoApi.PublishVideo)
			// 删除视频
			videoAuthGroup.POST("/delete", videoApi.DeleteVideo)
		}
	})
}
