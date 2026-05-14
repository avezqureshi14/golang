package payment

import (
	handler "payment/internal/payment/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(r *gin.Engine, h *handler.PaymentHandler) {
	payment := r.Group("/payment")
	payment.POST("/", h.CreatePayment)
}
