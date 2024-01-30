package dto

type LikeListDTO struct {
	UserID uint `json:"user_id" form:"user_id" binding:"required" message:"user_id cannot be empty"`
}

type LikeActionDTO struct {
	VideoID    uint `json:"video_id" form:"video_id" binding:"required" message:"video_id cannot be empty"`
	ActionType uint `json:"action_type" form:"action_type" binding:"required" message:"action_type cannot be empty"`
}
