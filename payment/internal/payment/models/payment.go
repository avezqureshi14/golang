package payment

type PaymentStatus string

const (
	StatusCreated  PaymentStatus = "CREATED"
	StatusApproved PaymentStatus = "APPROVED"
	StatusFailed   PaymentStatus = "FAILED"
)

type Payment struct {
	ID     int `gorm:"primaryKey"`
	UserID int
	Amount int
	Status PaymentStatus
}