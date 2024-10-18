package api

import (
	"context"
	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices/eventEmitter/internal/ports"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Application struct {
	event ports.EventPort
	rabbit    *amqp.Connection
}

func NewApplication(event ports.EventPort, rabbit *amqp.Connection) *Application {
	return &Application{
		event: event,
		rabbit: rabbit,
	}
}

func (a Application) AddLogEvent(ctx context.Context, eventEmitter domain.EventEmitter) (domain.EventEmitter, error) {
	err := a.event.LogEvent(ctx, &eventEmitter)
	if err != nil {
		return domain.EventEmitter{}, err
	}
	return eventEmitter, nil
}
