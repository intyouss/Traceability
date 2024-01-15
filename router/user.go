package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitUserRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup) {
		userApi := api.NewUserApi()
		userDefaultGroup := dfGroup.Group("user")
		{
			// 登录
			userDefaultGroup.POST("/login", userApi.Login)
			// 注册
			userDefaultGroup.POST("/register", userApi.Register)
		}
		userAuthGroup := auGroup.Group("user")
		{
			// 用户列表
			userAuthGroup.POST("/list", userApi.GetUserList)
			// 用户信息
			userAuthGroup.GET("/:id", userApi.GetUserInfo)
			// 更新用户
			userAuthGroup.PUT("/update", userApi.UpdateUser)
			// 删除用户
			userAuthGroup.DELETE("/:id", userApi.DeleteUser)
		}
	})
}
