package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup) {
		// 登录
		dfGroup.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "login Success",
			})
		})
		// 注册
		dfGroup.POST("/register", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(200, gin.H{
				"msg": "register success",
			})
		})

		userAuthGroup := auGroup.Group("user")
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

	})
}
