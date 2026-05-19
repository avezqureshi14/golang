package main

import (
	"log"
	"math/rand"
	"net/http"
	server "obs/internal/routes"
	"obs/pkg/telemetry"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Init tracing
	shutdown := telemetry.InitTracer()
	defer shutdown()

	// Init metrics
	telemetry.InitMetrics()

	// Setup router
	router := server.NewRouter()

	log.Println("Server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
