package ports

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
)

type APIPort interface {
	Charge(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error)
}
