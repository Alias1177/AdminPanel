package usecase

import (
	pb "AdminPanel/generated/go/proto"
	"AdminPanel/internal/domain/services"
	"context"
)

type Yield struct {
	yield services.YieldService
}

func NewYieldUseCase(yield services.YieldService) *Yield {
	return &Yield{yield: yield}
}

func (y *Yield) CheckCleanYield(ctx context.Context, req *pb.CheckYieldRequest) (*pb.CheckYieldResponse, error) {
	result, err := y.yield.GetCleanYield(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.CheckYieldResponse{
		Yield: result,
	}, nil
}
