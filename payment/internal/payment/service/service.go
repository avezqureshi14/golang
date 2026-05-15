package payment

import (
	dto "payment/internal/payment/dto"
	models "payment/internal/payment/models"
	repository "payment/internal/payment/repository"
)

type PaymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) Create(req *dto.CreatePaymentRequest) (dto.CreatePaymentResponse, error) {

	entity := &models.Payment{
		UserID: req.UserID,
		Amount: int(req.Amount),
		Status: models.StatusCreated,
	}

	created, err := s.repo.Create(entity)
	if err != nil {
		return dto.CreatePaymentResponse{}, err
	}

	return dto.CreatePaymentResponse{
		ID:     created.ID,
		UserID: created.UserID,
		Amount: created.Amount,
		Status: dto.PaymentStatus(created.Status),
	}, nil
}

func (s *PaymentService) GetAll() ([]dto.GetPaymentResponse, error) {

	payments, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	resp := make([]dto.GetPaymentResponse, 0, len(payments))

	for _, p := range payments {
		resp = append(resp, dto.GetPaymentResponse{
			ID:     p.ID,
			UserID: p.UserID,
			Amount: p.Amount,
			Status: dto.PaymentStatus(p.Status),
		})
	}

	return resp, nil
}

func (s *PaymentService) Update(id int, status dto.PaymentStatus) (dto.GetPaymentResponse, error) {

	updated, err := s.repo.Update(id, models.PaymentStatus(status))
	if err != nil {
		return dto.GetPaymentResponse{}, err
	}

	return dto.GetPaymentResponse{
		ID:     updated.ID,
		UserID: updated.UserID,
		Amount: updated.Amount,
		Status: dto.PaymentStatus(updated.Status),
	}, nil
}
