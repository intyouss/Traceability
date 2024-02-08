package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitLikeRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup) {
		likeApi := api.NewLikeApi()
		likeAuthGroup := auGroup.Group("like")
		{
			// 获取用户喜爱视频列表
			likeAuthGroup.GET("/list", likeApi.GetLikeList)
			// 用户喜爱操作
			likeAuthGroup.POST("/action", likeApi.LikeAction)
		}
	})
}
