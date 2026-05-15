package payment

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	dto "payment/internal/payment/dto"
	"testing"

	"github.com/gin-gonic/gin"
)

type PaymentServiceInterface interface {
	Create(*dto.CreatePaymentRequest) (dto.CreatePaymentResponse, error)
	GetAll() ([]dto.GetPaymentResponse, error)
	Update(int, dto.PaymentStatus) (dto.GetPaymentResponse, error)
}

type mockService struct {
	createFn func() (interface{}, error)
}

func TestCreatePayment_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	h := &PaymentHandler{}

	r := gin.Default()
	r.POST("/payment", h.CreatePayment)

	req := httptest.NewRequest(http.MethodPost, "/payment", bytes.NewBufferString(`invalid-json`))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}
