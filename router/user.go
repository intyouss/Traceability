package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
	"net/http"
)

func InitUserRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup) {
		userApi := api.NewUserApi()
		userDefaultGroup := dfGroup.Group("user")
		{
			// 登录
			userDefaultGroup.POST("/login", userApi.Login)
			// 注册
			userDefaultGroup.POST("/register", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(200, gin.H{
					"msg": "register success",
				})
			})
		}
		userAuthGroup := auGroup.Group("user")
		{
			// 用户列表
			userAuthGroup.GET("", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "zs"},
					},
				})
			})
			// 用户信息
			userAuthGroup.GET("/:id", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"id":   1,
					"name": "zs",
				})
			})
		}
	})
}
