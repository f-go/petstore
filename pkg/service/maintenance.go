package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type MaintenanceService struct {
}

// ...
func NewMaintenanceService() MaintenanceServiceServer {
	return &MaintenanceService{}
}

func (s *MaintenanceService) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *MaintenanceService) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return &EchoResponse{Feedback: "You sent: " + req.GetMessage()}, nil
}

