package dto

import (
	"github.com/intyouss/Traceability/models"
)

type Comment struct {
	ID        uint   `json:"id"`
	User      *User  `json:"user"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type CommentListDTO struct {
	CommonPageDTO
	VideoID uint `json:"video_id" form:"video_id" binding:"required" message:"video_id cannot be empty"`
}

// AddCommentDTO 添加评论数据传输对象
type AddCommentDTO struct {
	Content string `json:"content" form:"content" binding:"required" message:"content cannot be empty"`
	VideoId uint   `json:"video_id" form:"video_id" binding:"required" message:"video_id cannot be empty"`
}

func (c *AddCommentDTO) ToModel(comment *models.Comment) {
	comment.Content = c.Content
	comment.VideoId = c.VideoId
}

// DeleteCommentDTO 删除评论数据传输对象
type DeleteCommentDTO struct {
	CommonIDDTO
}
