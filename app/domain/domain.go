package domain

import (
	"time"
)

// Сообщение
type Message struct {
	Id         int       `json:"id" db:"id"`
	Text       string    `json:"text" db:"text"`
	ReceiverId int       `json:"receiver_id" db:"receiver_id"`
	SenderId   int       `json:"sender_id" db:"sender_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	IsRead     bool      `json:"is_read" db:"is_read"`
}

// Чат
type Thread struct {
	ReceiverId int `json:"receiver_id" db:"receiver_id"`
	SenderId   int `json:"sender_id" db:"sender_id"`
}
