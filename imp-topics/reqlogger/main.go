package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

/*
intercepting calls
observing data (status code)
forwarding to original implementation
*/
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// Middleware
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		// wrap response writer to capture status code
		rw := &responseWriter{
			ResponseWriter: w,
			status:         200, // default
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		log.Printf(
			"%s %s | status=%d | latency=%s",
			r.Method,
			r.URL.Path,
			rw.status,
			duration,
		)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	// wrap middleware
	loggedMux := RequestLogger(mux)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", loggedMux)
}
