package postgres

import (
    "log"
    "fmt"

    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations применяет миграции к базе данных PostgreSQL
func RunMigrations(dbHost, dbPort, dbUser, dbPassword, dbName string) error{
    // Строка подключения к базе данных
    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    // Создаём экземпляр migrate
    m, err := migrate.New(
        "file://migrations",
        connStr,
    )
    if err != nil {
        log.Fatalf("Failed to create migrate instance: %v", err)
		return err
    }

    // Применяем миграции
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Failed to run migrations: %v", err)
		return err
	}
    log.Println("Migrations applied successfully")
	return nil
}