package http

import (
	"net/http"
	"otus/messages/app/domain"
)

const (
	DB_ERROR            = 100
	PARAMETERS_REQUIRED = 101
)

// HTTP ответ
type Response struct {
	Code    int    `json:"code,omitempty"`    // код ответа (code != 0, если есть ошибка)
	Message string `json:"message,omitempty"` // текстовое сообщение
}

// Успешный ответ
func WithSuccess(message string) Response {
	return Response{
		Code:    http.StatusOK,
		Message: message,
	}
}

// Ответ с ошибкой
func WithError(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}

// Успешный ответ для выдачи сообщений
type MessagesResponse struct {
	Response
	Messages []domain.Message `json:"messages"` // сообщения
}

// Успешный ответ для выдачи количества непрочитанных
type GetNbUnreadResponse struct {
	Response
	Cnt int `json:"cnt"`
}
