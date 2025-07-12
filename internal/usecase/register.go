package usecase

import (
	pb "AdminPanel/generated/go/proto"
	"AdminPanel/internal/domain/entities"
	"context"
)

type RegisterAdmin struct {
	entities.AdminPanel
}

func (r *RegisterAdmin) RegisterAdmin(ctx context.Context, req *pb.RegReq) (*pb.RegResp, error) {
	return &pb.RegResp{}, nil
}
