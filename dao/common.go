package dao

import (
	"github.com/intyouss/Traceability/service/dto"
	"gorm.io/gorm"
)

// Paginate 通用分页函数
func Paginate(p dto.CommonPageDTO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := p.GetPage()
		limit := p.GetLimit()
		return db.Offset((page - 1) * limit).Limit(limit)
	}
}
