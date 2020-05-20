package database

import (
	"github.com/sirupsen/logrus"
	"otus/messages/app/domain"
	"otus/messages/app/usecase"
	"otus/messages/db"
	"time"
)

// Хранилище сообщений
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

	return
}

// Отметить как прочитанное
func (r *MessagesDatabase) MarkAsRead(receiverId int, senderId int) (err error) {
	sqlStatement := `UPDATE messages SET is_read=1 WHERE (receiver_id=? AND sender_id=?)`

	_, err = db.Connection().Exec(sqlStatement, receiverId, senderId)

	return
}

// Получение чатов
func (r *MessagesDatabase) GetThreads(userId int) (threads []domain.Thread, err error) {
	sqlStatement := `SELECT sender_id, receiver_id FROM messages WHERE sender_id=? 
		OR receiver_id=? GROUP BY sender_id, receiver_id`

	if err = db.Connection().Select(&threads, sqlStatement, userId, userId); err != nil {
		return
	}

	return
}

// Количество непрочитанных сообщений
func (r *MessagesDatabase) GetNbUnreadMessages(receiverId int) (cnt int, err error) {
	sqlStatement := `SELECT COUNT(*) AS count FROM messages WHERE receiver_id=? AND is_read=0`

	if err = db.Connection().Get(&cnt, sqlStatement, receiverId); err != nil {
		return
	}

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
