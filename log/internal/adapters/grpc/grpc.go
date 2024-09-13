package grpc

import (
	"context"
	"fmt"

	"github.com/ayuved/microservices-proto/golang/log"
	"github.com/ayuved/microservices/log/internal/application/core/domain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *log.CreateLogRequest) (*log.CreateLogResponse, error) {
	log.WithContext(ctx).Info("Creating log...")
	newLog := domain.NewLog(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newLog)
	log.WithContext(ctx).Infof("Log result: %v\n", result)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &log.CreateLogResponse{LogId: result.ID}, nil
}
