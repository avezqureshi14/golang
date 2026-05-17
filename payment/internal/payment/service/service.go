package payment

import (
	"context"
	"log"
	dto "payment/internal/payment/dto"
	models "payment/internal/payment/models"
	repository "payment/internal/payment/repository"
	fraud_client_grpc "payment/internal/platform/grpc"
	fraud_client_http "payment/internal/platform/http"
	"payment/pkg/jwt"

	"google.golang.org/grpc/metadata"
)

type PaymentService struct {
	repo                   repository.PaymentRepository
	fraudClientGRPCService *fraud_client_grpc.FraudClient
	fraudClientHTTPService *fraud_client_http.FraudHTTPClient
}

func NewPaymentService(repo repository.PaymentRepository, fraudClientGRPCService *fraud_client_grpc.FraudClient, fraudClientHTTPService *fraud_client_http.FraudHTTPClient) *PaymentService {
	return &PaymentService{repo: repo, fraudClientGRPCService: fraudClientGRPCService, fraudClientHTTPService: fraudClientHTTPService}
}

func (s *PaymentService) Create(ctx context.Context, req *dto.CreatePaymentRequest) (dto.CreatePaymentResponse, error) {

	// this was using shared api key
	// md := metadata.Pairs(
	// 	"x-api-key", "fraud-service-secret",
	// )

	// this one is using jwt token without sharing api secret
	token, err := jwt.GenerateServiceToken("payment-service")
	md := metadata.Pairs(
		"authorization", "Bearer "+token,
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx = metadata.NewOutgoingContext(context.Background(), md)
	fraudResp, err := s.fraudClientGRPCService.CheckFraud(ctx, int64(req.UserID), float64(req.Amount))
	// fraudResp, err := s.fraudClientHTTPService.CheckFraud(ctx, int64(req.UserID), float64(req.Amount))
	log.Print("my fraud response err", err)
	log.Print("my fraud response good", fraudResp)

	if err != nil {
		return dto.CreatePaymentResponse{}, err
	}

	var status models.PaymentStatus

	if fraudResp.IsFraud {
		status = models.StatusFailed
	} else {
		status = models.StatusApproved
	}

	entity := &models.Payment{
		UserID: req.UserID,
		Amount: int(req.Amount),
		Status: status,
	}

	created, err := s.repo.Create(entity)
	if err != nil {
		return dto.CreatePaymentResponse{}, err
	}

	return dto.CreatePaymentResponse{
		ID:     created.ID,
		UserID: created.UserID,
		Amount: created.Amount,
		Status: dto.PaymentStatus(created.Status),
	}, nil
}
func (s *PaymentService) GetAll() ([]dto.GetPaymentResponse, error) {

	payments, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	resp := make([]dto.GetPaymentResponse, 0, len(payments))

	for _, p := range payments {
		resp = append(resp, dto.GetPaymentResponse{
			ID:     p.ID,
			UserID: p.UserID,
			Amount: p.Amount,
			Status: dto.PaymentStatus(p.Status),
		})
	}

	return resp, nil
}

func (s *PaymentService) Update(id int, status dto.PaymentStatus) (dto.GetPaymentResponse, error) {

	updated, err := s.repo.Update(id, models.PaymentStatus(status))
	if err != nil {
		return dto.GetPaymentResponse{}, err
	}

	return dto.GetPaymentResponse{
		ID:     updated.ID,
		UserID: updated.UserID,
		Amount: updated.Amount,
		Status: dto.PaymentStatus(updated.Status),
	}, nil
}
