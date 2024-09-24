package ports

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
)

type DBPort interface {
	Get(ctx context.Context, id int64) (domain.Order, error)
	Save(context.Context, *domain.Order) error
}
