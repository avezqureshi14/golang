package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// connection pool config
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	return db
}