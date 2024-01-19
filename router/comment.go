package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitCommentRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup) {
		commentApi := api.NewCommentApi()
		commentDefaultGroup := dfGroup.Group("comment")
		{
			// 获取评论列表
			commentDefaultGroup.POST("/list", commentApi.GetCommentList)
		}
		commentAuthGroup := auGroup.Group("comment")
		{
			// 评论操作
			commentAuthGroup.POST("/action", commentApi.CommentAction)
		}
	})
}
