package usecase

import (
	pb "AdminPanel/generated/go/proto"
	"AdminPanel/internal/domain/entities"
	"context"
)

type Yield struct {
	entities.AdminPanel
}

func (y *Yield) CheckCleanYield(ctx context.Context, req *pb.CheckYieldRequest) (*pb.CheckYieldResponse, error) {
	return &pb.CheckYieldResponse{}, nil
}
