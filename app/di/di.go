package di

import (
	"github.com/gin-gonic/gin"
	"otus/messages/app/usecase"
)

// Инстанс приложения
type DI struct {
	MessagesDatabase usecase.MessagesDatabase
	CounterAPI       usecase.CounterAPI
}

// Новый инстанс приложения
func NewDI(messagesDatabase usecase.MessagesDatabase, counterAPI usecase.CounterAPI) DI {
	return DI{
		MessagesDatabase: messagesDatabase,
		CounterAPI:       counterAPI,
	}
}

// Пробрасывает зависимости в хэндлеры
func (di DI) ProvideDependency(f func(c *gin.Context, di DI)) func(*gin.Context) {
	return func(c *gin.Context) {
		f(c, di)
	}
}
