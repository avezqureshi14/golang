package middleware

import (
	"net/http"
	"obs/pkg/telemetry"
	"time"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &statusRecorder{ResponseWriter: w, status: 200}

		next.ServeHTTP(rec, r)

		duration := time.Since(start).Seconds()

		path := r.URL.Path
		method := r.Method
		status := http.StatusText(rec.status)

		telemetry.HttpRequestsTotal.WithLabelValues(path, method, status).Inc()
		telemetry.HttpDuration.WithLabelValues(path).Observe(duration)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}
