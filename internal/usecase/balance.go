package usecase

import (
	pb "AdminPanel/generated/go/proto"
	"AdminPanel/internal/domain/services"
	"fmt"

	"context"
)

type Balance struct {
	fetcher services.BalanceFetcher
}

func NewBalanceUseCase(fetcher services.BalanceFetcher) *Balance {
	return &Balance{fetcher: fetcher}
}
func (a *Balance) CheckTotalBalance(ctx context.Context, req *pb.CheckBalanceRequest) (*pb.CheckBalanceResponse, error) {
	balance, err := a.fetcher.GetBalance(ctx)
	if err != nil {
		return nil, fmt.Errorf("cant get balance: %w", err)
	}
	return &pb.CheckBalanceResponse{
		Balance: balance,
	}, nil
}
