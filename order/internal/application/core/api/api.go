package api // This package contains the API layer of the application

import (
	"github.com/ayuved/microservices/order/internal/application/core/domain"
	"github.com/ayuved/microservices/order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db}
}

func (a *Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)

	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
