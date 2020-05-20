package input

// Получение сообщений
type GetMessages struct {
	ReceiverId int `form:"receiverId" binding:"required"`
	SenderId   int `form:"senderId" binding:"required"`
}

// Получение чатов
type GetThreads struct {
	UserId int `form:"userId" binding:"required"`
}

// Создание сообщения
type CreateMessage struct {
	ReceiverId int    `form:"receiverId" binding:"required"`
	SenderId   int    `form:"senderId" binding:"required"`
	Text       string `form:"text" binding:"required"`
}
