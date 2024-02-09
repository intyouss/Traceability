package dto

type Message struct {
	// 源用户id
	FromUserID uint `json:"from_user_id"`
	// 目标用户id
	ToUserID uint `json:"to_user_id"`
	// 消息内容
	Content string `json:"content"`
	// 创建时间
	CreatedAt int64 `json:"created_at"`
}

// AddMessageDTO 添加消息数据传输对象
type AddMessageDTO struct {
	// 目标用户id
	ToUserID uint `json:"to_user_id" form:"to_user_id" binding:"required" message:"to_user_id cannot be empty"`
	// 消息内容
	Content string `json:"content" form:"content" binding:"required" message:"content cannot be empty"`
}

// MessageListDTO 消息列表数据传输对象
type MessageListDTO struct {
	// 目标用户id
	ToUserID uint `json:"to_user_id" form:"to_user_id" binding:"required" message:"to_user_id cannot be empty"`
	// 上一次消息时间
	PreMsgTime *int64 `json:"pre_msg_time" form:"pre_msg_time" binding:"required" message:"pre_msg_time cannot be empty"`
}
