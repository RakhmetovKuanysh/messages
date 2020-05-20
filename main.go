package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	nethttp "net/http"
	"os"
	"otus/messages/app/di"
	"otus/messages/app/http/router"
	"otus/messages/app/usecase/database"
	"otus/messages/counter_api"
	"otus/messages/db"
)

// конфиг микросервиса
var config struct {
	Listen   string
	Database db.Options
	Debug    bool
}

// инициализация конфига
func Init(configPath string) {
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		logrus.Fatalln("Не удалось загрузить конфиг", err)
	}
}

// точка входа
func main() {
	configPath := flag.String("config", "config/testing/config.toml", "Путь до файла конфига toml")
	flag.Parse()

	if *configPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	Init(*configPath)

	if err := db.Init(config.Database); err != nil {
		logrus.WithError(err).Fatalln("Не удалось инициализировать соединение к базе данных")
	}

	appDI := initDI()

	nethttp.Handle("/", router.Router(appDI, config.Debug))

	logrus.Info("ms-messages - сервис запущен")
	logrus.Fatal(nethttp.ListenAndServe(config.Listen, nil))
}

// инициализация приложения
func initDI() *di.DI {
	counterAPI := counter_api.NewCounterAPI()
	messagesDatabase := database.NewMessagesDatabase()
	core := di.NewDI(messagesDatabase, counterAPI)

	return &core
}
