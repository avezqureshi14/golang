package fraud

import models "fraud/models"

type FraudResponse struct {
	IsFraud bool   `json:"isFraud"`
	Message string `json:"message"`
}

func CheckFraud(p models.Payment) FraudResponse {
	// Dummy rules (keep deterministic for benchmarking)
	if p.Amount > 100000 {
		return FraudResponse{
			IsFraud: true,
			Message: "High amount transaction - Fraud detected",
		}
	}

	if p.Status == models.StatusFailed {
		return FraudResponse{
			IsFraud: true,
			Message: "Failed transaction pattern - Fraud suspected",
		}
	}

	return FraudResponse{
		IsFraud: false,
		Message: "Transaction is safe",
	}
}
