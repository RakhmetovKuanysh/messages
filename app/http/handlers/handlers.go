package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"otus/messages/app/di"
	"otus/messages/app/domain"
	app "otus/messages/app/http"
	"otus/messages/app/http/input"
)

// Проверка состояния сервиса
func Health(c *gin.Context, di di.DI) {
	c.String(http.StatusOK, "ok")
}

// Получение всех семестров
func GetMessages(c *gin.Context, di di.DI) {
	in := input.GetMessages{}

	if err := c.MustBindWith(&in, binding.Query); err != nil {
		return
	}

	if in.ReceiverId == 0 || in.SenderId == 0 {
		c.JSON(http.StatusBadRequest, app.WithError(app.PARAMETERS_REQUIRED, "Provide parameters"))

		return
	}

	messages, err := di.MessagesDatabase.GetMessages(in.ReceiverId, in.SenderId)

	if err != nil {
		c.JSON(http.StatusBadRequest, app.WithError(app.DB_ERROR, "DB Error"))

		return
	}

	if messages == nil {
		messages = make([]domain.Message, 0)
	}

	c.JSON(http.StatusOK, app.MessagesResponse{
		Response: app.WithSuccess("Found"),
		Messages: messages,
	})

	return
}

// Добавление сообщения
func CreateMessage(c *gin.Context, di di.DI) {
	in := input.CreateMessage{}

	if err := c.MustBindWith(&in, binding.Form); err != nil {
		return
	}

	if in.ReceiverId == 0 || in.SenderId == 0 || in.Text == "" {
		c.JSON(http.StatusBadRequest, app.WithError(app.PARAMETERS_REQUIRED, "Provide parameters"))

		return
	}

	err := di.MessagesDatabase.CreateMessage(in.ReceiverId, in.SenderId, in.Text)

	if err != nil {
		c.JSON(http.StatusBadRequest, app.WithError(app.DB_ERROR, "DB Error"))

		return
	}

	c.JSON(http.StatusOK, app.WithSuccess("Added"))

	return
}
