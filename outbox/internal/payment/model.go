package payment

import "time"

type Payment struct {
	ID        string
	UserID    string
	Amount    float64
	Status    string
	CreatedAt time.Time
}