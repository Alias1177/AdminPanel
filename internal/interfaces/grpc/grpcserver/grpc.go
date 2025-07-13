package grpcserver

import (
	pb "AdminPanel/generated/go/proto"
	"google.golang.org/grpc"
)

func NewGRPCServer(handler pb.AdminPanelServer) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterAdminPanelServer(server, handler)
	return server
}
