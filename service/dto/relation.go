package dto

type RelationActionDto struct {
	// 关注者id
	UserID uint `json:"user_id" form:"user_id" binding:"required" message:"user_id cannot be empty"`
	// 动作种类
	ActionType uint `json:"action_type" form:"action_type" binding:"required" message:"action_type cannot be empty"`
}

type FocusListDto struct {
	// 用户id
	UserID uint `json:"user_id" form:"user_id" binding:"required" message:"user_id cannot be empty"`
	CommonPageDTO
}

type FansListDto struct {
	// 用户id
	UserID uint `json:"user_id" form:"user_id" binding:"required" message:"user_id cannot be empty"`
	CommonPageDTO
}
