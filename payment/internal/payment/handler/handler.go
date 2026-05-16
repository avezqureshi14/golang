package payment

import (
	"fmt"
	"net/http"
	dto "payment/internal/payment/dto"
	service "payment/internal/payment/service"
	http_response "payment/internal/platform/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req dto.CreatePaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		http_response.SendError(c, http.StatusBadRequest, "invalid body")
		return
	}

	res, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		http_response.HandleError(c, err)
		return
	}

	http_response.SendSuccess(c, http.StatusCreated, res)
}

func (h *PaymentHandler) GetAllPayments(c *gin.Context) {

	res, err := h.service.GetAll()
	if err != nil {
		http_response.HandleError(c, err)
		return
	}

	http_response.SendSuccess(c, http.StatusOK, res)
}

func (h *PaymentHandler) UpdatePayment(c *gin.Context) {

	idParam := c.Param("id")

	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		http_response.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req dto.UpdatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_response.SendError(c, http.StatusBadRequest, "invalid body")
		return
	}

	res, err := h.service.Update(id, req.Status)
	if err != nil {
		http_response.HandleError(c, err)
		return
	}

	http_response.SendSuccess(c, http.StatusOK, res)
}
