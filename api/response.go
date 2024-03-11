package api

import (
	"net/http"
	"reflect"

	"github.com/intyouss/Traceability/global"

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
	global.Logger.Infof(
		"Request: (Method) %v | (Path) %v | (Form) %v | (Code) %v | (Msg) %v | (Data) %+v",
		ctx.Request.Method, ctx.Request.RequestURI, ctx.Request.Form, resp.Code, resp.Msg, resp.Data,
	)
}

// Fail 失败响应
func Fail(ctx *gin.Context, resp *Response) {
	HttpResponse(ctx, resp)
	global.Logger.Errorf(
		"Request: (Method) %v | (Path) %v | (Form) %v | (ErrorCode) %v | (ErrorMsg) %v",
		ctx.Request.Method, ctx.Request.RequestURI, ctx.Request.Form, resp.Code, resp.Msg,
	)
}
