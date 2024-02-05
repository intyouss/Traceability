package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitRelationRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup) {
		relationApi := api.NewRelationApi()
		//relationDefaultGroup := dfGroup.Group("relation")
		//{
		//
		//}
		relationAuthGroup := auGroup.Group("relation")
		{
			// 关注/取消关注
			relationAuthGroup.POST("/action", relationApi.RelationAction)
			// 关注列表
			relationAuthGroup.GET("/focus/list", relationApi.GetFocusList)
			// 粉丝列表
			relationAuthGroup.GET("/fans/list", relationApi.GetFansList)
		}
	})
}
