package ports

import (
	"context"
	"github.com/ayuved/microservices-helper/domain"
)

type DBPort interface {
	Get(ctx context.Context, id string) (domain.Payment, error)
	Save(ctx context.Context, payment *domain.Payment) error
}