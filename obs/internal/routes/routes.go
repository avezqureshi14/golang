package server

import (
	"net/http"
	"obs/internal/handler"
	"obs/pkg/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	var hello http.Handler = http.HandlerFunc(handler.HelloHandler)

	hello = otelhttp.NewHandler(hello, "hello-endpoint")
	hello = middleware.MetricsMiddleware(hello)

	mux.Handle("/api/hello", hello)
	mux.Handle("/metrics", promhttp.Handler())

	return mux
}
