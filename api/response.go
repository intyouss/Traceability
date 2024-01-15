package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Response struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  int64  `json:"total,omitempty"`
}

// IsEmpty 判断响应实体是否为空
func (r *Response) IsEmpty() bool {
	return reflect.DeepEqual(r, Response{})
}

// HttpResponse 自定义响应
func HttpResponse(ctx *gin.Context, status int, resp *Response) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	ctx.AbortWithStatusJSON(status, resp)
}

// setStatus 设置响应状态码
func setStatus(resp *Response, status int) int {
	if resp.Status == 0 {
		return status
	}
	return resp.Status
}

// Success 成功响应
func Success(ctx *gin.Context, resp *Response) {
	HttpResponse(ctx, setStatus(resp, http.StatusOK), resp)
}

// Fail 失败响应
func Fail(ctx *gin.Context, resp *Response) {
	HttpResponse(ctx, setStatus(resp, http.StatusBadRequest), resp)
}

// ServerError 服务器错误响应
func ServerError(ctx *gin.Context, resp *Response) {
	HttpResponse(ctx, setStatus(resp, http.StatusInternalServerError), resp)
}
