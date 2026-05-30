package outbox

import (
	"database/sql"
	"encoding/json"
)

func handlePayment(db *sql.DB, payload string) {
	var data map[string]interface{}
	json.Unmarshal([]byte(payload), &data)

	paymentID := data["payment_id"]

	db.Exec(`
		UPDATE payments
		SET status='SUCCESS'
		WHERE id=$1
	`, paymentID)
}

func process(db *sql.DB) {
	rows, _ := db.Query(`
		SELECT id, event_type, payload
		FROM outbox_events
		WHERE processed = false
		LIMIT 50
	`)

	defer rows.Close()

	for rows.Next() {
		var id, eventType string
		var payload string

		rows.Scan(&id, &eventType, &payload)

		if eventType == "PAYMENT_CREATED" {
			handlePayment(db, payload)
		}

		db.Exec(`UPDATE outbox_events SET processed=true WHERE id=$1`, id)
	}
}