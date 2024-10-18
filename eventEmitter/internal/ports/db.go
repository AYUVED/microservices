package ports

import (
	"context"
	"github.com/ayuved/microservices-helper/domain"
)

type EventPort interface {
	LogEvent(ctx context.Context, eventEmitter *domain.EventEmitter) error
}
