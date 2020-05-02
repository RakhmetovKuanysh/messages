package input

// Отправка сообщения
type GetMessages struct {
	ReceiverId int `form:"receiverId" binding:"required"`
	SenderId   int `form:"senderId" binding:"required"`
}

// Создание сообщения
type CreateMessage struct {
	ReceiverId int    `form:"receiverId" binding:"required"`
	SenderId   int    `form:"senderId" binding:"required"`
	Text       string `form:"text" binding:"required"`
}
