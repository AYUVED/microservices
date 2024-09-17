package ports

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
)

type APIPort interface {
	Add(ctx context.Context, logservice domain.Logservice) (domain.Logservice, error)
	Get(ctx context.Context, id string) (domain.Logservice, error)
}
