package global

import (
	"github.com/intyouss/Traceability/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger *zap.SugaredLogger
	DB     *gorm.DB
	Redis  *config.RedisClient
)
