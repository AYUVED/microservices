package ports

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
