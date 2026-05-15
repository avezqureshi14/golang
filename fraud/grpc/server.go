package grpcserver

import (
	"context"

	"fraud/fraud/proto"
	models "fraud/models"
	service "fraud/service"
)

type Server struct {
	proto.UnimplementedFraudServiceServer
}

func (s *Server) CheckFraud(ctx context.Context, req *proto.PaymentRequest) (*proto.FraudResponse, error) {
	payment := models.Payment{
		ID:     int(req.Id),
		UserID: int(req.UserId),
		Amount: int(req.Amount),
		Status: models.PaymentStatus(req.Status),
	}

	result := service.CheckFraud(payment)

	return &proto.FraudResponse{
		IsFraud: result.IsFraud,
		Message: result.Message,
	}, nil
}
