package dto

import (
	"github.com/intyouss/Traceability/models"
)

type CommentListDTO struct {
	CommonPageDTO
	VideoID uint `json:"video_id" form:"video_id" binding:"required" message:"VideoId cannot be empty"`
}

// CommentActionDTO 评论操作数据传输对象
type CommentActionDTO struct {
	Content    string `json:"content,omitempty" form:"content"`
	UserID     uint   `json:"user_id" form:"user_id" binding:"required" message:"UserId cannot be empty"`
	VideoID    uint   `json:"video_id" form:"video_id" binding:"required" message:"VideoId cannot be empty"`
	ActionType uint   `json:"action_type" form:"action_type" binding:"required" message:"ActionType cannot be empty"`
	CommentID  uint   `json:"comment_id,omitempty" form:"comment_id"`
}

func (c *CommentActionDTO) AddActionToModel(comment *models.Comment) {
	comment.Content = c.Content
	comment.UserId = c.UserID
	comment.VideoId = c.VideoID
}
