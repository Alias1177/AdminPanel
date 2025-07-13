package grpc

import (
	"AdminPanel/internal/domain/services"
	"context"

	pb "AdminPanel/generated/go/proto"
)

type AdminGRPCHandler struct {
	pb.UnimplementedAdminPanelServer
	balanceUC services.BalanceUseCase
	loginUC   services.LoginUseCase
	driverUC  services.DriverUseCase
	client    services.ClientUseCase
	yieldUC   services.YieldUseCase
}

func NewAdminGRPCHandler(bUC services.BalanceUseCase, loginUC services.LoginUseCase, driverUC services.DriverUseCase, client services.ClientUseCase, yieldUC services.YieldUseCase) *AdminGRPCHandler {
	return &AdminGRPCHandler{balanceUC: bUC, loginUC: loginUC, driverUC: driverUC, client: client, yieldUC: yieldUC}
}

func (h *AdminGRPCHandler) CheckTotalBalance(ctx context.Context, req *pb.CheckBalanceRequest) (*pb.CheckBalanceResponse, error) {
	return h.balanceUC.CheckTotalBalance(ctx, req)
}
func (h *AdminGRPCHandler) LoginAdmin(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return h.loginUC.LoginAdmin(ctx, req)
}
func (h *AdminGRPCHandler) CheckDriverData(ctx context.Context, req *pb.DataDriverRequest) (*pb.DataDriverResponse, error) {
	return h.driverUC.CheckDriverData(ctx, req)
}
func (h *AdminGRPCHandler) CheckClientData(ctx context.Context, req *pb.DataClientRequest) (*pb.DataClientResponse, error) {
	return h.client.CheckClientData(ctx, req)
}
func (h *AdminGRPCHandler) CheckCleanYield(ctx context.Context, req *pb.CheckYieldRequest) (*pb.CheckYieldResponse, error) {
	return h.yieldUC.CheckCleanYield(ctx, req)
}
func (h *AdminGRPCHandler) GetCleanYield(ctx context.Context, req *pb.CheckYieldRequest) (*pb.CheckYieldResponse, error) {
	return h.yieldUC.CheckCleanYield(ctx, req)
}
