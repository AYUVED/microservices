package api

import (
	"context"
	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices/shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error) {
	err := a.db.Save(ctx, &shipping)
	if err != nil {
		return domain.Shipping{}, err
	}
	return shipping, nil
}