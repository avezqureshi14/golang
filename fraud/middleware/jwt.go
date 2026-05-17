package middleware

import (
	"context"
	"strings"

	auth "fraud/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func JWTAuthInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}

	tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")

	claims, err := auth.ValidateToken(tokenStr)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	// optional: inject identity into context
	ctx = context.WithValue(ctx, "service", claims.Service)

	return handler(ctx, req)
}
