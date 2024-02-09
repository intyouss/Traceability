package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	SuccessCode = 0
	SuccessMsg  = "success"
)

type Response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg,omitempty"`
	Data  any    `json:"data,omitempty"`
	Total int64  `json:"total,omitempty"`
}

// IsEmpty 判断响应实体是否为空
func (r *Response) IsEmpty() bool {
	return reflect.DeepEqual(r, Response{})
}

// HttpResponse 自定义响应
func HttpResponse(ctx *gin.Context, resp *Response) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(http.StatusOK)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

// Success 成功响应
func Success(ctx *gin.Context, resp *Response) {
	resp.Code = SuccessCode
	resp.Msg = SuccessMsg
	HttpResponse(ctx, resp)
}

// Fail 失败响应
func Fail(ctx *gin.Context, resp *Response) {
	HttpResponse(ctx, resp)
}
