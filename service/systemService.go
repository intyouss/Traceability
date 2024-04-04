package service

import (
	"context"
	"fmt"

	"github.com/intyouss/Traceability/utils"
)

var SystemServiceIns *SystemService

type SystemService struct {
	BaseService *BaseService
}

func NewSystemService() *SystemService {
	if SystemServiceIns == nil {
		SystemServiceIns = &SystemService{
			BaseService: NewBaseService(),
		}
	}
	return SystemServiceIns
}

// GetCpuUsage 获取CPU使用率
func (s *SystemService) GetCpuUsage(ctx context.Context) (string, error) {
	usage, err := utils.GetTotalCpuUsage(ctx)
	if err != nil {
		return "", err
	}
	//
	a := fmt.Sprintf("%.2f%%", usage)
	return a, nil
}

// GetMemoryUsage 获取内存使用率
func (s *SystemService) GetMemoryUsage(ctx context.Context) (string, error) {
	usage, err := utils.GetMemoryUsage(ctx)
	if err != nil {
		return "", err
	}
	a := fmt.Sprintf("%.2f%%", usage)
	return a, nil
}
