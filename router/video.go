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
		}
		videoAuthGroup := auGroup.Group("video")
		{
			//// 视频操作
			//videoAuthGroup.POST("/action", videoApi.VideoAction)
			// 用户发布视频列表
			videoAuthGroup.GET("/list", videoApi.GetUserVideoList)
		}
	})
}
