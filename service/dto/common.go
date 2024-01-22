package dto

// CommonUserIDDTO 通用ID数据传输对象
type CommonUserIDDTO struct {
	ID uint `json:"user_id" form:"user_id"`
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
