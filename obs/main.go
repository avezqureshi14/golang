package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 🔥 Metrics
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method"},
	)

	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpDuration)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// simulate work
	time.Sleep(200 * time.Millisecond)

	w.Write([]byte("hello world"))

	duration := time.Since(start).Seconds()

	httpRequestsTotal.WithLabelValues("/hello", r.Method).Inc()
	httpDuration.WithLabelValues("/hello").Observe(duration)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	// Prometheus endpoint
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
