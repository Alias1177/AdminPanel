package usecase

import (
	pb "AdminPanel/generated/go/proto"
	"AdminPanel/internal/domain/services"
	"context"
)

type Client struct {
	client services.ClientService
}

func NewClientUseCase(client services.ClientService) *Client {
	return &Client{client: client}
}

func (c *Client) CheckClientData(ctx context.Context, req *pb.DataClientRequest) (*pb.DataClientResponse, error) {
	client, err := c.client.GetClientData(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.DataClientResponse{
		ClientInfo: client,
	}, nil
}
