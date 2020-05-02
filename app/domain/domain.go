package domain

import (
	"time"
)

// Сообщение
type Message struct {
	Id         int       `json:"id" db:"id"`
	Text       string    `json:"text" db:"text"`
	ReceiverId int       `json:"receiverId" db:"receiver_id"`
	SenderId   int       `json:"senderId" db:"sender_id"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	IsRead     bool      `json:"isRead" db:"is_read"`
}
