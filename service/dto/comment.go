package dto

import (
	"github.com/intyouss/Traceability/models"
)

type Comment struct {
	ID       uint   `json:"id"`
	User     *User  `json:"user"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
}

type CommentListDTO struct {
	CommonPageDTO
	VideoID uint `json:"video_id" form:"video_id" binding:"required" message:"video_id cannot be empty"`
}

// AddCommentDTO 添加评论数据传输对象
type AddCommentDTO struct {
	Content string `json:"content" form:"content"`
	UserID  uint   `json:"id" form:"id" binding:"required" message:"user_id cannot be empty"`
	VideoID uint   `json:"video_id" form:"video_id" binding:"required" message:"video_id cannot be empty"`
}

func (c *AddCommentDTO) ToModel(comment *models.Comment) {
	comment.Content = c.Content
	comment.UserId = c.UserID
	comment.VideoId = c.VideoID
}

// DeleteCommentDTO 删除评论数据传输对象
type DeleteCommentDTO struct {
	UserID    uint `json:"id" form:"id" binding:"required" message:"user_id cannot be empty"`
	CommentID uint `json:"comment_id" form:"comment_id" binding:"required" message:"comment_id cannot be empty"`
}
