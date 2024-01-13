package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// Login 登录Api
// @Summary 用户登录
// @Description 用户登录
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} json "{"msg": "success"}"
// @Failure 401 {string} json "{"msg": "error"}"
// @Router /api/v1/public/user/login [post]
func (u UserApi) Login(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "login",
	})
}
