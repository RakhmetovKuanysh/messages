package usecase

import "otus/messages/app/domain"

// Репозиторий отчетов
type MessagesDatabase interface {
	CreateMessage(receiverId int, senderId int, text string) (err error)
	GetMessages(receiverId int, senderId int) (messages []domain.Message, err error)
}
