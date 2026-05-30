package main

import (
	"log"
	"net/http"

	"outbox/db"
	"outbox/internal/api"
	"outbox/internal/cron"
	"outbox/internal/outbox"
	"outbox/internal/payment"
)

func main() {
	dburl := "postgresql://postgres.xgkgtnwfcpnwtelkkjmo:golang9890562214@aws-1-ap-southeast-1.pooler.supabase.com:6543/postgres?sslmode=require&pgbouncer=true"

	conn, err := db.Connect(dburl)
	if err != nil {
		log.Fatal(err)
	}

	// services
	paymentService := payment.NewService(conn)
	handler := api.NewHandler(paymentService)

	// worker + cron
	outbox.StartWorker(conn)
	cron.StartCron(conn)

	http.HandleFunc("/payment", handler.CreatePayment)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}