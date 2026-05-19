package telemetry

import "github.com/prometheus/client_golang/prometheus"

var (
	HttpRequestsTotal *prometheus.CounterVec
	HttpDuration      *prometheus.HistogramVec
)

func InitMetrics() {
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)

	HttpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	prometheus.MustRegister(HttpRequestsTotal, HttpDuration)
}
