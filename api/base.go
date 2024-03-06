package api

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/intyouss/Traceability/global"
	"go.uber.org/zap"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() *BaseApi {
	return &BaseApi{
		Logger: global.Logger,
	}
}

// BuildRequestOption 构建请求参数选项
type BuildRequestOption struct {
	Ctx *gin.Context
	DTO any
}

// BuildRequest 构建请求参数
func (b *BaseApi) BuildRequest(opt BuildRequestOption) *BaseApi {
	var errResult error
	// 绑定请求上下文
	b.Ctx = opt.Ctx
	// 绑定请求参数
	if opt.DTO != nil {
		err := b.Ctx.ShouldBind(opt.DTO)
		if err != nil {
			if errResult == nil {
				errResult = err
			} else {
				errResult = fmt.Errorf("%w, %w", errResult, err)
			}
		}
		if errResult != nil {
			errResult = b.parseValidateError(errResult, opt.DTO)
			b.AddError(errResult)
		}
	}
	return b
}

func (b *BaseApi) AddError(err error) {
	if b.Errors == nil {
		b.Errors = err
		return
	}
	b.Errors = fmt.Errorf("%w, %w", b.Errors, err)
}

func (b *BaseApi) GetError() error {
	return b.Errors
}

// parseValidateError 解析并修改校验错误提示
func (b *BaseApi) parseValidateError(errs error, target any) (errResult error) {
	var errValidation validator.ValidationErrors
	if ok := errors.As(errs, &errValidation); !ok {
		errResult = errs
		return
	}

	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}
		if errResult == nil {
			errResult = errors.New(errMessage)
			continue
		}
		errResult = fmt.Errorf("%w, %w", errResult, errors.New(errMessage))
	}
	return
}

func (b *BaseApi) Fail(resp *Response) {
	Fail(b.Ctx, resp)
}

func (b *BaseApi) Success(resp *Response) {
	Success(b.Ctx, resp)
}
