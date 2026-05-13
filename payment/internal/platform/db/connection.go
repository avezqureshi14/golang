package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// Get underlying sql.DB to configure pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB:", err)
	}

	// connection pool config
	
	// Limit total number of open connections to DB
	// Use case:
	// Prevents overwhelming the database under high traffic
	// Extra requests will wait instead of opening new connections
	sqlDB.SetMaxOpenConns(25)

	// Keep up to 10 idle (unused) connections ready
	// Use case:
	// Avoids cost of creating new DB connections for every request
	// Improves performance by reusing existing connections
	sqlDB.SetMaxIdleConns(10)

	// Maximum lifetime of a connection (here: 1 hour)
	// Use case:
	// Prevents using stale/broken connections (due to network, DB restarts, etc.)
	// Forces periodic refresh of connections for stability
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
