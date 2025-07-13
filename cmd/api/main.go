package main

import (
	"AdminPanel/config"
	grpc "AdminPanel/internal/interfaces/grpc/fetch"
	"AdminPanel/internal/interfaces/grpc/grpcserver"
	"AdminPanel/internal/interfaces/http/handlers"
	"AdminPanel/internal/usecase"
	"log"
	"net"
)

func main() {
	cfg, err := config.Load()

	balanceUC := usecase.NewBalanceUseCase(&handlers.GetBalance{
		ApiURL: cfg.ApiUrl,
	})

	loginFetcher := usecase.NewLoginUseCase(&handlers.LoginUs{
		ApiUrl: cfg.ApiUrl,
	})

	driverData := usecase.NewDriverUseCase(&handlers.DriverReq{
		ApiUrl: cfg.ApiUrl,
	})

	clientData := usecase.NewClientUseCase(&handlers.ClientReq{
		ApiUrl: cfg.ApiUrl,
	})

	adminHandler := grpc.NewAdminGRPCHandler(balanceUC, loginFetcher, driverData, clientData)
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
