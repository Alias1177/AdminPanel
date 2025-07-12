package main

import (
	"AdminPanel/config"
	"AdminPanel/internal/interfaces/grpc/grpcserver"
	grpc "AdminPanel/internal/interfaces/grpc/handlers"
	"AdminPanel/internal/interfaces/http/handlers"
	"AdminPanel/internal/usecase"
	"log"
	"net"
)

func main() {
	cfg, err := config.Load()
	balanceFetcher := &handlers.GetBalance{
		ApiURL: cfg.ApiUrl,
	}
	balanceUC := usecase.NewBalanceUseCase(balanceFetcher)
	adminHandler := grpc.NewAdminGRPCHandler(balanceUC)

	grpcServer := grpcserver.NewGRPCServer(adminHandler)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
