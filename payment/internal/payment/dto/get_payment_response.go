package payment

type GetPaymentResponse struct {
	ID     int           `json:"id"`
	UserID int           `json:"userId"`
	Amount int           `json:"amount"`
	Status PaymentStatus `json:"status"`
}
