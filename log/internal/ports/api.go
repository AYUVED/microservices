package ports

import (
	"context"

	"github.com/ayuved/microservices/log/internal/application/core/domain"
)

type APIPort interface {
	LogItem(ctx context.Context, log domain.Log) (domain.Log, error)
}
