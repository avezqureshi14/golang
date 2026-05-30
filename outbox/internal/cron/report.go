package cron

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func generateReport(db *sql.DB) {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	rows, _ := db.Query(`
		SELECT id, user_id, amount, created_at
		FROM payments
		WHERE status='SUCCESS'
		AND DATE(created_at) = $1
	`, yesterday)

	defer rows.Close()

	var result []map[string]interface{}

	for rows.Next() {
		var id, userID string
		var amount float64
		var createdAt time.Time

		rows.Scan(&id, &userID, &amount, &createdAt)

		result = append(result, map[string]interface{}{
			"id": id,
			"user_id": userID,
			"amount": amount,
			"created_at": createdAt,
		})
	}

	fileName := fmt.Sprintf("payments-%s.json", yesterday)
	file, _ := os.Create(fileName)
	defer file.Close()

	json.NewEncoder(file).Encode(result)
}

func StartCron(db *sql.DB) {
	c := cron.New()

	c.AddFunc("0 9 * * *", func() {
		generateReport(db)
	})

	c.Start()
}