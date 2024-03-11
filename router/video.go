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
			// 主页视频feed流
			videoDefaultGroup.GET("/feed", videoApi.GetVideoFeed)
			// 用户发布视频列表
			videoDefaultGroup.GET("/list", videoApi.GetUserVideoList)
			// 视频搜索
			videoDefaultGroup.GET("/search", videoApi.GetVideoSearch)
		}
		videoAuthGroup := auGroup.Group("video")
		{
			// 推荐，关注，朋友页视频feed流
			videoAuthGroup.GET("/feed", videoApi.GetVideoFeed)
			// 用户发布视频列表
			videoAuthGroup.GET("/list", videoApi.GetUserVideoList)
			// 发布视频
			videoAuthGroup.POST("/publish", videoApi.PublishVideo)
			// 删除视频
			videoAuthGroup.POST("/delete", videoApi.DeleteVideo)
			// 获取单个视频
			videoAuthGroup.GET("/info", videoApi.GetVideoInfo)
		}
	})
}
