package store

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open() (*sql.DB, error) {
	specifications := "host=localhost user=postgres password=postgres password=postgres dbname=postgres port=5433 sslmode=disabled"

	db, err := sql.Open("pgx", specifications)

	if err != nil {
		return nil, fmt.Errorf("db open: %w", err)
	}

	fmt.Printf("Connecting to the Database ...\n")
	return db, nil
}
