package postgres

import (
	"errors"
	"fmt"
	
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(host, port, user, password, dbname string) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=&s port=%s user=%s password=%s dbname=%s sslmode=disable",
	 host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		err = errors.New("не удалось подключиться к базе данных")
		slog.Error("Ошибка подключения базы данных", err, "connStr", connStr)
		return nil, err
	}
	return db, nil
}