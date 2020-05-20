package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"net/http"
	"otus/messages/app/di"
	"otus/messages/app/domain"
	app "otus/messages/app/http"
	"otus/messages/app/http/input"
	"strconv"
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
		c.JSON(http.StatusBadRequest, app.WithError(app.DB_ERROR, "DB Error while getting messages"))
		return
	}

	err = di.MessagesDatabase.MarkAsRead(in.ReceiverId, in.SenderId)

	if err != nil {
		c.JSON(http.StatusBadRequest, app.WithError(app.DB_ERROR, "DB Error while marking as read"))
		return
	}

	go di.CounterAPI.UnsetNbUnread(in.ReceiverId)

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

	go updateUnreadMessages(di, in.ReceiverId)

	c.JSON(http.StatusOK, app.WithSuccess("Added"))

	return
}

// Получает количество непрочитанных сообщений
func GetNbUnread(c *gin.Context, di di.DI) {
	userId, ok := c.GetQuery("userId")

	if !ok {
		c.JSON(http.StatusBadRequest, app.WithError(app.PARAMETERS_REQUIRED, "Provide parameters"))
		return
	}

	receiverid, err := strconv.Atoi(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, app.WithError(app.PARAMETERS_REQUIRED, "Provide valid parameters"))
		return
	}

	nbCnt, err := di.MessagesDatabase.GetNbUnreadMessages(receiverid)

	if err != nil {
		logrus.WithField("receiverid", receiverid).
			WithError(err).
			Error("Не удалось получить количество непрочитанных сообщений")
		return
	}

	c.JSON(http.StatusOK, app.GetNbUnreadResponse{
		Response: app.WithSuccess("Found"),
		Cnt:      nbCnt,
	})

	return
}

// Обновляет количестве непрочитанных сообщений в фоновом режиме
func updateUnreadMessages(di di.DI, receiverId int) {
	nbCnt, err := di.MessagesDatabase.GetNbUnreadMessages(receiverId)

	if err != nil {
		logrus.WithField("receiverid", receiverId).
			WithError(err).
			Error("Не удалось получить количество непрочитанных сообщений")
		return
	}

	di.CounterAPI.SetNbUnread(receiverId, nbCnt)
}
