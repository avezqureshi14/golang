package middleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const APIKeyHeader = "x-api-key"
const ValidAPIKey = "fraud-service-secret"

func APIKeyAuthInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	keys := md[APIKeyHeader]
	log.Print("My fraud service", keys)
	if len(keys) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing api key")
	}

	if keys[0] != ValidAPIKey {
		return nil, status.Error(codes.Unauthenticated, "invalid api key")
	}

	// continue request flow
	return handler(ctx, req)
}
