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
