package outbox

import (
	"database/sql"
	"time"
)

func StartWorker(db *sql.DB) {
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		for range ticker.C {
			process(db)
		}
	}()
}