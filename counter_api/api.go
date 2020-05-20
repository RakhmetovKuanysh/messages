package counter_api

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"otus/messages/app/usecase"
	"strconv"
)

// API
type CounterAPI struct {
}

// Новый клиент
func NewCounterAPI() usecase.CounterAPI {
	return &CounterAPI{}
}

// Устанавливает количество непрочитанных сообщений
func (r *CounterAPI) SetNbUnread(userId int, cnt int) (err error) {
	formData := url.Values{
		"userId": {strconv.Itoa(userId)},
		"cnt":    {strconv.Itoa(cnt)},
	}

	resp, err := http.PostForm("http://127.0.0.1:8090/set-nb-unread", formData)

	if err != nil {
		logrus.WithError(err).Error("Ошибка при обновлении непрочитанных сообщений")
		return
	}

	defer resp.Body.Close()

	return
}

// Сбрасывает количество непрочитанных сообщений
func (r *CounterAPI) UnsetNbUnread(userId int) (err error) {
	resp, err := http.Get("http://127.0.0.1:8090/unset-nb-unread?userId=" + strconv.Itoa(userId))

	if err != nil {
		logrus.WithError(err).Error("Ошибка при простроении запроса для получения количества непрочитанных сообщений")
		return
	}

	defer resp.Body.Close()

	return
}
