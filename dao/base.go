package dao

import (
	"github.com/intyouss/Traceability/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BaseDao struct {
	DB     *gorm.DB
	logger *zap.SugaredLogger
}

func NewBaseDao() *BaseDao {
	return &BaseDao{
		DB:     global.DB,
		logger: global.Logger,
	}
}
