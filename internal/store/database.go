package store

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Open() (*sql.DB, error) {
	specifications := "host=localhost user=postgres password=postgres  dbname=postgres port=5433 sslmode=disable"

	db, err := sql.Open("pgx", specifications)

	if err != nil {
		return nil, fmt.Errorf("db open: %w", err)
	}

	fmt.Printf("Connecting to the Database ...\n")
	return db, nil
}

func MigrateFS(db *sql.DB, migrateFS fs.FS, dir string) error {
	goose.SetBaseFS(migrateFS)

	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")

	if err != nil {
		return fmt.Errorf("migrate error:%w", err)
	}

	err = goose.Up(db, dir)

	if err != nil {
		return fmt.Errorf("goose Up: %w", err)
	}
	return nil
}
