package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(dbURL string) (*sql.DB, error) {
	return sql.Open("postgres", dbURL)
}