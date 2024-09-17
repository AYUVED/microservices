package ports

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
)

type LogServicePort interface {
	PlaceOrder(ctx context.Context, o *domain.Order) error
	GetOrder(ctx context.Context, id int64) (domain.Order, error)
}
