package global

import (
	"github.com/intyouss/Traceability/config"
	"github.com/intyouss/Traceability/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger *zap.SugaredLogger
	DB     *gorm.DB
	Redis  *config.RedisClient
	OSS    *utils.MinioClient
)
