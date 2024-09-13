package api

import (
	"context"

	"github.com/ayuved/microservices/log/internal/application/core/domain"
	"github.com/ayuved/microservices/log/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(ctx context.Context, log domain.Log) (domain.Log, error) {
	err := a.db.Save(ctx, &log)
	if err != nil {
		return domain.Log{}, err
	}
	return log, nil
}
