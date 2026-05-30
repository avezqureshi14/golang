package api

import (
	"encoding/json"
	"net/http"

	"outbox/internal/payment"

	"github.com/google/uuid"
)

type Handler struct {
	Service *payment.Service
}

func NewHandler(s *payment.Service) *Handler {
	return &Handler{Service: s}
}

type PaymentRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

func (h *Handler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	json.NewDecoder(r.Body).Decode(&req)

	id := uuid.New().String()

	err := h.Service.CreatePayment(id, req.UserID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"payment_id": id,
		"status":     "PENDING",
	})
}