package api

import (
	"github.com/gin-gonic/gin"
	"github.com/intyouss/Traceability/service"
)

const (
	ErrGetCpuUsage = iota + 80001
	ErrGetMemoryUsage
)

type SystemApi struct {
	BaseApi
	Service *service.SystemService
}

func NewSystemApi() SystemApi {
	return SystemApi{
		BaseApi: NewBaseApi(),
		Service: service.NewSystemService(),
	}
}

// GetCpuUsage
// @Summary 获取CPU使用率
// @Param token header string true "token"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /api/v1/admin/system/cpu/usage [get]
func (s *SystemApi) GetCpuUsage(c *gin.Context) {
	if err := s.BuildRequest(BuildRequestOption{Ctx: c}).GetError(); err != nil {
		s.Fail(&Response{Code: ErrGetCpuUsage, Msg: err.Error()})
		return
	}
	usage, err := s.Service.GetCpuUsage(c)
	if err != nil {
		s.Fail(&Response{Code: ErrGetCpuUsage, Msg: err.Error()})
		return
	}
	s.Success(&Response{Data: gin.H{"cpu_usage": usage}})
}

// GetMemoryUsage
// @Summary 获取内存使用率
// @Param token header string true "token"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /api/v1/admin/system/memory/usage [get]
func (s *SystemApi) GetMemoryUsage(ctx *gin.Context) {
	if err := s.BuildRequest(BuildRequestOption{Ctx: ctx}).GetError(); err != nil {
		s.Fail(&Response{Code: ErrGetMemoryUsage, Msg: err.Error()})
		return
	}
	usage, err := s.Service.GetMemoryUsage(ctx)
	if err != nil {
		s.Fail(&Response{Code: ErrGetMemoryUsage, Msg: err.Error()})
		return
	}
	s.Success(&Response{
		Data: gin.H{
			"memory_usage": usage,
		},
	})
}
