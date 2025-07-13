package usecase

import (
	pb "AdminPanel/generated/go/proto"
	"AdminPanel/internal/domain/services"
	"context"
)

type Driver struct {
	drive services.DriverService
}

func NewDriverUseCase(drive services.DriverService) *Driver {
	return &Driver{drive: drive}
}

func (d *Driver) CheckDriverData(ctx context.Context, req *pb.DataDriverRequest) (*pb.DataDriverResponse, error) {
	drive, err := d.drive.GetDriverData(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.DataDriverResponse{
		DriverInfo: drive,
	}, nil
}
