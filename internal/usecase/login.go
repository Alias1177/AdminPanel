package usecase

import (
	pb "AdminPanel/generated/go/proto"
	"AdminPanel/internal/domain/services"
	"context"
)

type Login struct {
	loginService services.LoginService
}

func NewLoginUseCase(loginService services.LoginService) *Login {
	return &Login{loginService: loginService}
}

func (l *Login) LoginAdmin(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	login, err := l.loginService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		Response: login,
	}, nil
}
