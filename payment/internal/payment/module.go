package payment

import (
	payment_handler "payment/internal/payment/handler"
	payment_repo "payment/internal/payment/repository"
	payment_routes "payment/internal/payment/routes"
	payment_service "payment/internal/payment/service"
	fraud_client_http "payment/internal/platform/http"

	fraud_client_grpc "payment/internal/platform/grpc"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	h *payment_handler.PaymentHandler
}

func New(db *gorm.DB) *Module {
	repo := payment_repo.NewPaymentRepo(db)
	fraudClientGRPCService, _ := fraud_client_grpc.NewFraudClient("localhost:50051")
	fraudClientHTTPService := fraud_client_http.NewFraudHTTPClient("http://localhost:8081")
	service := payment_service.NewPaymentService(repo, fraudClientGRPCService, fraudClientHTTPService)
	handler := payment_handler.NewPaymentHandler(service)

	return &Module{h: handler}
}

func (m *Module) RegisterRoutes(r *gin.Engine) {
	payment_routes.RegisterPaymentRoutes(r, m.h)
}
