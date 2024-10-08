package ports

import (
	"context"
	"github.com/ayuved/microservices-helper/domain"
)

type DBPort interface {
	Get(ctx context.Context, id string) (domain.Shipping, error)
	Save(ctx context.Context, shipping *domain.Shipping) error
}