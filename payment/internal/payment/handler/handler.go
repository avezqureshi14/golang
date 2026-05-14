package payment

import (
	"net/http"
	dto "payment/internal/payment/dto"
	service "payment/internal/payment/service"
	http_response "payment/internal/platform/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: *service,
	}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req dto.CreatePaymentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		http_response.SendError(c, http.StatusBadRequest, "invalid body")
		return
	}
	payment, err := h.service.Create(&dto.CreatePaymentRequest{UserID: req.UserID, Amount: req.Amount})
	if err != nil {
		http_response.HandleError(c, err)
		return
	}
	http_response.SendSuccess(c, http.StatusCreated, payment)
}
