package payment

import "database/sql"

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) CreatePayment(id, userID string, amount float64) error {
	_, err := s.DB.Exec(`
		SELECT create_payment($1, $2, $3)
	`, id, userID, amount)

	return err
}