package service

import (
	"github.com/intyouss/Traceability/global"
	"go.uber.org/zap"
)

type BaseService struct {
	logger *zap.SugaredLogger
}

func NewBaseService() *BaseService {
	return &BaseService{
		logger: global.Logger,
	}
}
