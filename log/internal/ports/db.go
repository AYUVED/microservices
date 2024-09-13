package ports

import (
	"context"

	"github.com/ayuved/microservices/log/internal/application/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id string) (domain.Log, error)
	LogItem(ctx context.Context, log *domain.Log) error
}
