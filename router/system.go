package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitSystemRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup) {
		systemApi := api.NewSystemApi()
		systemAuthGroup := auGroup.Group("system")
		{
			// 获取CPU使用率
			systemAuthGroup.GET("/cpu/usage", systemApi.GetCpuUsage)
			// 获取内存使用率
			systemAuthGroup.GET("/memory/usage", systemApi.GetMemoryUsage)
		}
	})
}
