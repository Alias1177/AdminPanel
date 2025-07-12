// internal/interfaces/grpc/admin_handler.go
package grpc

import (
	"AdminPanel/internal/domain/services"
	"context"

	pb "AdminPanel/generated/go/proto"
)

type AdminGRPCHandler struct {
	pb.UnimplementedAdminPanelServer
	balanceUC services.BalanceUseCase
}

func NewAdminGRPCHandler(bUC services.BalanceUseCase) *AdminGRPCHandler {
	return &AdminGRPCHandler{balanceUC: bUC}
}

func (h *AdminGRPCHandler) CheckTotalBalance(ctx context.Context, req *pb.CheckBalanceRequest) (*pb.CheckBalanceResponse, error) {
	return h.balanceUC.CheckTotalBalance(ctx, req)
}
