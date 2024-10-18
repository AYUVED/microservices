package grpc

import (
	"context"
	"fmt"

	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices-proto/golang/eventEmitter"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *eventEmitter.CreateLogEventRequest) (*eventEmitter.CreateLogEventResponse, error) {
	log.WithContext(ctx).Info("Creating eventEmitter...")
	neweventEmitter := domain.NewEventEmitter(request.App,request.Name, request.Data, request.ProcessId, request.Type,request.Status, request.User)
	result, err := a.api.AddLogEvent(ctx, neweventEmitter)
	log.WithContext(ctx).Infof("eventEmitter result: %v\n", result)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &eventEmitter.CreateLogEventResponse{ 
		Id: result.ID,
	 }, nil
}
