package grpc

import (
	"context"
	"fmt"

	"github.com/ayuved/microservices-proto/golang/logservice"
	"github.com/ayuved/microservices-helper/domain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Add(ctx context.Context, request *logservice.CreateLogRequest) (*logservice.CreateLogResponse, error) {
	log.WithContext(ctx).Info("Creating logservice...")
	log.Println("Creating logservice...", request)
	newLogservice := domain.NewLogservice(request.App, request.Name, request.Data)
	result, err := a.api.Add(ctx, newLogservice)

	log.WithContext(ctx).Infof("Logservice result: %v\n", result)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &logservice.CreateLogResponse{Id: result.ID}, nil
}
