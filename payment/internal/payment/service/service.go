package payment

import (
	dto "payment/internal/payment/dto"
	payment "payment/internal/payment/models"
	repository "payment/internal/payment/repository"
)

type PaymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) *PaymentService{
	return &PaymentService{
		repo: repo,
	}
}

func (s *PaymentService) Create (p * dto.CreatePaymentRequest) error{
	err := s.repo.Create(&payment.Payment{UserID: p.UserID,Amount: int(p.Amount),Status: payment.StatusCreated})
	if err != nil{
		return err
	}
	return nil
}