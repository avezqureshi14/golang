package payment

import (
	"errors"
	"testing"

	dto "payment/internal/payment/dto"
	models "payment/internal/payment/models"
)

type mockRepo struct {
	createFn func(*models.Payment) (models.Payment, error)
	getAllFn func() ([]models.Payment, error)
	updateFn func(int, models.PaymentStatus) (models.Payment, error)
}

func (m *mockRepo) Create(p *models.Payment) (models.Payment, error) {
	return m.createFn(p)
}

func (m *mockRepo) GetAll() ([]models.Payment, error) {
	return m.getAllFn()
}

func (m *mockRepo) Update(id int, status models.PaymentStatus) (models.Payment, error) {
	return m.updateFn(id, status)
}

func TestPaymentService_Create(t *testing.T) {
	mock := &mockRepo{
		createFn: func(p *models.Payment) (models.Payment, error) {
			p.ID = 1
			return *p, nil
		},
	}

	svc := NewPaymentService(mock)

	req := &dto.CreatePaymentRequest{
		UserID: 1,
		Amount: 500,
	}

	res, err := svc.Create(req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if res.ID != 1 {
		t.Errorf("expected id 1, got %d", res.ID)
	}
}

func TestPaymentService_Create_Error(t *testing.T) {
	mock := &mockRepo{
		createFn: func(p *models.Payment) (models.Payment, error) {
			return models.Payment{}, errors.New("db error")
		},
	}

	svc := NewPaymentService(mock)

	req := &dto.CreatePaymentRequest{
		UserID: 1,
		Amount: 100,
	}

	_, err := svc.Create(req)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestPaymentService_Update(t *testing.T) {
	mock := &mockRepo{
		updateFn: func(id int, status models.PaymentStatus) (models.Payment, error) {
			return models.Payment{
				ID:     id,
				Status: status,
			}, nil
		},
	}

	svc := NewPaymentService(mock)

	res, err := svc.Update(1, dto.PaymentStatus("SUCCESS"))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.Status != dto.PaymentStatus("SUCCESS") {
		t.Errorf("status mismatch")
	}
}
