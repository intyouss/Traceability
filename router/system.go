package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitSystemRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup, adGroup *gin.RouterGroup) {
		systemApi := api.NewSystemApi()
		systemAdminGroup := adGroup.Group("system")
		// 后台系统功能
		{
			// 获取CPU使用率
			systemAdminGroup.GET("/cpu/usage", systemApi.GetCpuUsage)
			// 获取内存使用率
			systemAdminGroup.GET("/memory/usage", systemApi.GetMemoryUsage)
		}
	})
}
