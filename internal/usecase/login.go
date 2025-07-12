package usecase

import (
	pb "AdminPanel/generated/go/proto"
	"AdminPanel/internal/domain/entities"
	"context"
)

type LoginAdmin struct {
	entities.AdminPanel
}

func (l *LoginAdmin) LoginAdmin(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {

	return &pb.LoginResp{}, nil
}
