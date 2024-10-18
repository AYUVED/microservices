package ports

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
)

type APIPort interface {
	AddLogEvent(ctx context.Context, eventEmitter domain.EventEmitter) (domain.EventEmitter, error)
}
