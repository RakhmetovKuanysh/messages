package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Параметры подключения
type Options struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

var DB *sqlx.DB

// Открывает соединение с бд
func Init(options Options) error {
	conn, err := sqlx.Connect("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		options.User,
		options.Password,
		options.Host,
		options.Port,
		options.Database,
	))

	DB = conn

	return err
}

// Пингует соединение с бд
func PingDb() error {
	return DB.Ping()
}

// Получить подключение
func Connection() *sqlx.DB {
	return DB
}
