package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/api"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/models"
	"github.com/intyouss/Traceability/utils"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ErrCodeInvalidToken = iota + 10401
	ErrCodeTokenParse
	ErrCodeTokenNotMatched
	ErrCodeTokenRenew
	TokenKey    = "Authorization"
	TokenPrefix = "Bearer: "
)

func tokenError(c *gin.Context, code int) {
	api.Fail(c, &api.Response{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    "invalid token",
	})
}

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader(TokenKey)

		var id string
		switch c.Request.Method {
		case http.MethodGet:
			id = c.Query("id")
			if id == "" {
				id = c.Param("id")
			}
		default:
			id = c.PostForm("id")
		}

		if token == "" || !strings.HasPrefix(token, TokenPrefix) {
			tokenError(c, ErrCodeInvalidToken)
			return
		}

		token = strings.TrimPrefix(token, TokenPrefix)
		claims, err := utils.ParseToken(token)
		if err != nil {
			tokenError(c, ErrCodeTokenParse)
			return
		}

		// 检查用户id是否一致
		userId := strconv.Itoa(int(claims.ID))
		if userId != id {
			tokenError(c, ErrCodeTokenNotMatched)
			return
		}

		if claims.ExpiresAt.Time.Before(time.Now().Add(
			-viper.GetDuration("jwt.tokenExpire")*time.Minute + 20*time.Minute),
		) {
			newToken, err := utils.GenerateToken(claims.ID, claims.Name)
			if err != nil {
				tokenError(c, ErrCodeTokenRenew)
				return
			}
			c.Header("token", newToken)
		}

		c.Set(global.LoginUser, models.LoginUser{
			ID:       claims.ID,
			Username: claims.Name,
		})
		c.Next()
	}
}
