package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitializeMigrations(fileString, connString string) error {
	m, err := migrate.New(fileString, connString)
	if err != nil {
		return err
	}

	return m.Up()
}
