package payment

type PaymentStatus string

const (
	StatusCreated  PaymentStatus = "CREATED"
	StatusApproved PaymentStatus = "APPROVED"
	StatusFailed   PaymentStatus = "FAILED"
)

type Payment struct {
	ID     int           `gorm:"primaryKey" json:"id"`
	UserID int           `json:"userId"`
	Amount int           `json:"amount"`
	Status PaymentStatus `json:"status"`
}
