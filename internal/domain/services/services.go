package services

import (
	pb "AdminPanel/generated/go/proto"
	"context"
)

type BalanceFetcher interface {
	GetBalance(ctx context.Context) (string, error)
}
type BalanceUseCase interface {
	CheckTotalBalance(ctx context.Context, req *pb.CheckBalanceRequest) (*pb.CheckBalanceResponse, error)
}

type LoginService interface {
	Login(ctx context.Context, email, password string) (string, error)
}

type LoginUseCase interface {
	LoginAdmin(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error)
}

type DriverService interface {
	GetDriverData(ctx context.Context) (string, error)
}
type DriverUseCase interface {
	CheckDriverData(ctx context.Context, req *pb.DataDriverRequest) (*pb.DataDriverResponse, error)
}
type ClientService interface {
	GetClientData(ctx context.Context) (string, error)
}
type ClientUseCase interface {
	CheckClientData(ctx context.Context, req *pb.DataClientRequest) (*pb.DataClientResponse, error)
}
