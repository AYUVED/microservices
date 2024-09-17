package ports

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
)

type DBPort interface {
	Get(ctx context.Context, id string) (domain.Logservice, error)
	Add(ctx context.Context, logservice *domain.Logservice) error
}
