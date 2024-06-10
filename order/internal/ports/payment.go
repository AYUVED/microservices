package ports

import (
	"context"

	"github.com/ayuved/microservices/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
