package dto

type Message struct {
	// 源用户
	FromUser *User `json:"from_user"`
	// 目标用户
	ToUser *User `json:"to_user"`
	// 消息内容
	Content string `json:"content"`
	// 创建时间
	CreatedAt int64 `json:"created_at"`
}

type MessageOpen struct {
	// 开放联系人
	OpenUser *User `json:"open_user"`
	// 消息列表
	Messages []*Message `json:"messages"`
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
	PreMsgTime string `json:"pre_msg_time" form:"pre_msg_time" binding:"required" message:"pre_msg_time cannot be empty"`
}

type LinkUsersDTO struct {
	UserId uint `json:"user_id" form:"user_id" binding:"required" message:"user_id cannot be empty"`
}

type OpenMsgListDTO struct {
	UserId uint `json:"user_id" form:"user_id" binding:"required" message:"user_id cannot be empty"`
}

type AddOpenUserDTO struct {
	OpenUserID uint `json:"open_user_id" form:"open_user_id" binding:"required" message:"open_user_id cannot be empty"`
}

type DeleteOpenUserDTO struct {
	OpenUserID uint `json:"open_user_id" form:"open_user_id" binding:"required" message:"open_user_id cannot be empty"`
}
