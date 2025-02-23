package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // драйвер PostgreSQL
)

// NewDB устанавливает соединение с PostgreSQL и возвращает объект *sqlx.DB.
func NewDB(host, port, user, password, dbname string) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return sqlx.Connect("postgres", connStr)
}
