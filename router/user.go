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
			// 用户信息
			userDefaultGroup.GET("/", userApi.GetUserInfo)
			// 模糊搜索用户列表
			userDefaultGroup.GET("/search", userApi.GetUserListBySearch)
		}
		userAuthGroup := auGroup.Group("user")
		{
			// 用户信息
			userAuthGroup.GET("/", userApi.GetUserInfo)
			// 更新用户
			userAuthGroup.POST("/update", userApi.UpdateUser)
			// 删除用户
			userAuthGroup.POST("/delete", userApi.DeleteUser)
			// 模糊搜索用户列表
			userAuthGroup.GET("/search", userApi.GetUserListBySearch)
		}
	})
}
