package payment

import models "payment/internal/payment/models"

type PaymentRepository interface {
	Create(p *models.Payment) error
	GetAll() ([]models.Payment, error)
	Update(id int, status models.PaymentStatus) (models.Payment, error)
}
