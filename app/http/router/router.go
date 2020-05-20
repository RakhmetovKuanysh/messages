package router

import (
	"github.com/gin-gonic/gin"
	"otus/messages/app/di"
	"otus/messages/app/http/handlers"
)

// Инициализирует маршрутизатор
func Router(p *di.DI, debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	if debug {
		r.Use(gin.Logger())
	}

	r.Use(Recovery())

	configureRoutes(r, p)

	return r
}

// Настройки маршрутов
func configureRoutes(r *gin.Engine, p *di.DI) {
	r.GET("/health", p.ProvideDependency(handlers.Health))
	r.GET("/messages", p.ProvideDependency(handlers.GetMessages))
	r.POST("/message", p.ProvideDependency(handlers.CreateMessage))
	r.GET("/get-nb-unread", p.ProvideDependency(handlers.GetNbUnread))
}
