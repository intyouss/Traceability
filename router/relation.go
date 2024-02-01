package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitRelationRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup) {
		relationApi := api.NewRelationApi()
		relationDefaultGroup := dfGroup.Group("relation")
		{
			// 关注列表
			relationDefaultGroup.GET("/focus/list", relationApi.GetFocusList)
			// 粉丝列表
			relationDefaultGroup.GET("/fans/list", relationApi.GetFansList)
		}
		relationAuthGroup := auGroup.Group("relation")
		{
			// 关注/取消关注
			relationAuthGroup.POST("/action", relationApi.RelationAction)

		}
	})
}
