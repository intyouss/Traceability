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
			// 添加评论
			commentAuthGroup.POST("/add", commentApi.AddComment)
			// 删除评论
			commentAuthGroup.POST("/delete", commentApi.DeleteComment)
		}
	})
}
