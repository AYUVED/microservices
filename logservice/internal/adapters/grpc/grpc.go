package grpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices-proto/golang/logservice"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Add(ctx context.Context, request *logservice.CreateLogRequest) (*logservice.CreateLogResponse, error) {
	log.WithContext(ctx).Info("Creating logservice...")
	var data interface{}
	log.Printf("Request123: %v\n", request.Data)
	err := json.Unmarshal([]byte(request.Data), &data)
    if err != nil {
        log.Fatalf("Error unmarshaling JSON: %v", err)
    }
	newLogservice := domain.NewLogservice(request.App, request.Name,
		data, request.ProcessId, request.Status, request.Type, request.User)
	result, err := a.api.Add(ctx, newLogservice)

	log.WithContext(ctx).Infof("Logservice result: %v\n", result)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &logservice.CreateLogResponse{Id: result.ID}, nil
}
