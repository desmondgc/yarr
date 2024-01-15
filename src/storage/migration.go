package storage

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func migrate(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

    if err := goose.SetDialect("postgres"); err != nil {
        return err
    }

    if err := goose.Up(db, "migrations"); err != nil {
        return err
    }

	return nil
}
