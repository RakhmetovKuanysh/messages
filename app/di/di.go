package di

import (
	"github.com/gin-gonic/gin"
	"otus/messages/app/usecase"
)

// Инстанс приложения
type DI struct {
	MessagesDatabase usecase.MessagesDatabase
}

// Новый инстанс приложения
func NewDI(messagesDatabase usecase.MessagesDatabase) DI {
	return DI{
		MessagesDatabase: messagesDatabase,
	}
}

// Пробрасывает зависимости в хэндлеры
func (di DI) ProvideDependency(f func(c *gin.Context, di DI)) func(*gin.Context) {
	return func(c *gin.Context) {
		f(c, di)
	}
}
