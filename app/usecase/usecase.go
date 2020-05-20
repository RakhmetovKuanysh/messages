package usecase

import "otus/messages/app/domain"

// Репозиторий сообщений
type MessagesDatabase interface {
	CreateMessage(receiverId int, senderId int, text string) (err error)
	GetMessages(receiverId int, senderId int) (messages []domain.Message, err error)
	GetNbUnreadMessages(receiverId int) (cnt int, err error)
	MarkAsRead(receiverId int, senderId int) (err error)
}

// Клиент для работы с CounterAPI
type CounterAPI interface {
	SetNbUnread(userId int, cnt int) (err error)
	UnsetNbUnread(userId int) (err error)
}
