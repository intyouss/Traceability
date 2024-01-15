package dao

import (
	"github.com/intyouss/Traceability/global"
	"gorm.io/gorm"
)

type BaseDao struct {
	DB *gorm.DB
}

func NewBaseDao() *BaseDao {
	return &BaseDao{
		DB: global.DB,
	}
}
