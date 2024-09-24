package api

import (
	"context"
	"log"

	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices/logservice/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Add(ctx context.Context, logservice domain.Logservice) (domain.Logservice, error) {
	log.Println("Adding logservice", logservice)
	err := a.db.Add(ctx, &logservice)
	if err != nil {
		return domain.Logservice{}, err
	}
	return logservice, nil
}

func (a Application) Get(ctx context.Context, id string) (domain.Logservice, error) {
	l, err := a.db.Get(ctx, id)
	if err != nil {
		return domain.Logservice{}, err
	}
	return l, nil
}
