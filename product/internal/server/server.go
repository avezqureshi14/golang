package server

import (
	"net/http"

	handler "product/internal/handler"
)

func Start(h *handler.Handler) {
	http.HandleFunc("/view", h.ViewHandler)
	http.HandleFunc("/product", h.GetProductHandler)

	http.ListenAndServe(":8080", nil)
}