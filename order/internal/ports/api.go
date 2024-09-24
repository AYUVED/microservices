package ports

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	GetOrder(ctx context.Context, id int64) (domain.Order, error)
}
