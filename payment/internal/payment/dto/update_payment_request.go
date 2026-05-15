package payment

type UpdatePaymentRequest struct {
	Status PaymentStatus `json:"status"`
}
