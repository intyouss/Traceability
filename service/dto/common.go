package dto

// CommonIDDTO 通用ID数据传输对象
type CommonIDDTO struct {
	ID *uint `json:"id" form:"id" binding:"required" message:"id cannot be empty"`
}

// CommonPageDTO 通用分页数据传输对象
type CommonPageDTO struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

func (c *CommonPageDTO) GetPage() int {
	if c.Page <= 0 {
		return 1
	}
	return c.Page
}

func (c *CommonPageDTO) GetLimit() int {
	if c.Limit <= 0 {
		return 10
	}
	return c.Limit
}
