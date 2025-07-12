package usecase

import (
	pb "AdminPanel/generated/go/proto"
	"AdminPanel/internal/domain/entities"
	"context"
)

type Driver struct {
	entities.AdminPanel
}

func (d *Driver) CheckDriverData(ctx context.Context, req *pb.DataDriverRequest) (*pb.DataDriverResponse, error) {
	return &pb.DataDriverResponse{}, nil
}
