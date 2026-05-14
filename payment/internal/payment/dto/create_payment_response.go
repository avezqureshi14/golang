package payment

type PaymentStatus string

const (
	StatusCreated  PaymentStatus = "CREATED"
	StatusApproved PaymentStatus = "APPROVED"
	StatusFailed   PaymentStatus = "FAILED"
)

type CreatePaymentResponse struct {
	ID     int
	UserID int
	Amount int
	Status PaymentStatus
}
