package router

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
)

func InitUserRoutes() {
	RegisterRoute(func(dfGroup *gin.RouterGroup, auGroup *gin.RouterGroup, adGroup *gin.RouterGroup) {
		userApi := api.NewUserApi()
		userDefaultGroup := dfGroup.Group("user")
		{
			// 登录
			userDefaultGroup.POST("/login", userApi.Login)
			// 注册
			userDefaultGroup.POST("/register", userApi.Register)
			// 用户信息
			userDefaultGroup.GET("/", userApi.GetUserInfo)
			// 获取用户列表
			userDefaultGroup.GET("/list", userApi.GetUserList)
		}
		userAuthGroup := auGroup.Group("user")
		{
			// 用户信息
			userAuthGroup.GET("/", userApi.GetUserInfo)
			// 更新用户
			userAuthGroup.POST("/update", userApi.UpdateUser)
			// 获取用户列表
			userAuthGroup.GET("/list", userApi.GetUserList)
			// 上传头像
			userAuthGroup.POST("/upload/avatar", userApi.UploadAvatar)
			// 取消上传头像
			userAuthGroup.POST("/upload/avatar/abolish", userApi.AbolishAvatarUpload)

		}
		// 后台系统功能
		userAdminGroup := adGroup.Group("user")
		{
			// 删除用户
			userAdminGroup.POST("/delete", userApi.DeleteUser)
			// 获取月总日用户增长记录
			userAdminGroup.GET("/increase", userApi.GetUserIncrease)
			// 获取角色列表
			userAdminGroup.GET("/role/list", userApi.GetRoleList)
			// 添加角色
			userAdminGroup.POST("/role/add", userApi.AddRole)
			// 删除角色
			userAdminGroup.POST("/role/delete", userApi.DeleteRole)
			// 更新角色
			userAdminGroup.POST("/role/update", userApi.UpdateRole)
			// 获取用户列表
			userAdminGroup.GET("/list", userApi.GetUserList)
			// 更新用户
			userAdminGroup.POST("/update", userApi.UpdateUser)
			// 获取用户总数
			userAdminGroup.GET("/total", userApi.GetUserTotal)
			// 添加用户
			userAdminGroup.POST("/add", userApi.Register)
		}
	})
}
