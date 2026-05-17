package main

import (
	"log"
	"net"
	"net/http"

	"fraud/fraud/proto"
	grpcserver "fraud/grpc"
	handler "fraud/handler"
	"fraud/middleware"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// 🔹 Start Gin (HTTP)
	go func() {
		r := gin.Default()
		handler.RegisterRoutes(r)

		log.Println("HTTP server running on :8081")
		http.ListenAndServe(":8081", r)
	}()

	// 🔹 Start gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.APIKeyAuthInterceptor),
	)
	proto.RegisterFraudServiceServer(grpcServer, &grpcserver.Server{})

	log.Println("gRPC server running on :50051")
	grpcServer.Serve(lis)
}
