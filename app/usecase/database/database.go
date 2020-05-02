package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"otus/messages/app/domain"
	"otus/messages/app/usecase"
	"otus/messages/db"
	"time"
)

// Хранилище отчетов
type MessagesDatabase struct {
}

// Новое хранилище сообщений
func NewMessagesDatabase() usecase.MessagesDatabase {
	return &MessagesDatabase{}
}

// Получение сообщений
func (r *MessagesDatabase) GetMessages(receiverId int, senderId int) (messages []domain.Message, err error) {
	sqlStatement := `SELECT * FROM messages WHERE (receiver_id=? AND sender_id=?)
		OR (sender_id=? AND receiver_id=?) ORDER BY created_at`

	err = db.Connection().Select(&messages, sqlStatement, receiverId, senderId, receiverId, senderId)
	fmt.Println(err)

	return
}

// Создание сообщения
func (r *MessagesDatabase) CreateMessage(receiverId int, senderId int, text string) (err error) {
	sqlStatement := `INSERT INTO messages(sender_id, receiver_id, text, created_at) VALUES (?, ?, ?, ?) `

	_, err = db.Connection().Exec(sqlStatement, senderId, receiverId, text, time.Now())

	if err != nil {
		logrus.WithError(err).Error("Ошибка запроса на добавление сообщения в бд")
		return
	}

	return
}
