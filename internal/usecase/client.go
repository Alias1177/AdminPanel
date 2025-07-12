package usecase

import (
	pb "AdminPanel/generated/go/proto"
	"AdminPanel/internal/domain/entities"
	"context"
)

type Client struct {
	entities.AdminPanel
}

func (c *Client) CheckClientData(ctx context.Context, req *pb.DataClientRequest) (*pb.DataClientResponse, error) {
	return &pb.DataClientResponse{}, nil
}
