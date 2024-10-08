package ports

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
