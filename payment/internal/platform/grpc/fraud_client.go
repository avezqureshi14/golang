package payment

import (
	"context"
	"time"

	fraud_proto "payment/internal/proto/fraud/proto"

	"google.golang.org/grpc"
)

type FraudClient struct {
	conn   *grpc.ClientConn
	client fraud_proto.FraudServiceClient
}

func NewFraudClient(addr string) (*FraudClient, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(), // replace with TLS in prod
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	client := fraud_proto.NewFraudServiceClient(conn)

	return &FraudClient{
		conn:   conn,
		client: client,
	}, nil
}

func (f *FraudClient) Close() {
	f.conn.Close()
}

/*
to control and carry the lifecycle of a request across multiple layers and services.
Without it, your system becomes “fire-and-forget chaos”.
*/
func (f *FraudClient) CheckFraud(ctx context.Context, userID int64, amount float64) (*fraud_proto.FraudResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	resp, err := f.client.CheckFraud(ctx, &fraud_proto.PaymentRequest{
		UserId: int32(userID),
		Amount: int32(amount),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
