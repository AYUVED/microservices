package db

import (
	"context"
	"fmt"

	"github.com/ayuved/microservices/log/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"

	//"gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderID    int64
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Log, error) {
	var logEntity Log
	res := a.db.WithContext(ctx).First(&logEntity, id)
	log := domain.Log{
		ID:         int64(logEntity.ID),
		CustomerID: logEntity.CustomerID,
		Status:     logEntity.Status,
		OrderId:    logEntity.OrderID,
		TotalPrice: logEntity.TotalPrice,
		CreatedAt:  logEntity.CreatedAt.UnixNano(),
	}
	return log, res.Error
}

func (a Adapter) Save(ctx context.Context, log *domain.Log) error {
	orderModel := Log{
		CustomerID: log.CustomerID,
		Status:     log.Status,
		OrderID:    log.OrderId,
		TotalPrice: log.TotalPrice,
	}
	res := a.db.WithContext(ctx).Create(&orderModel)
	if res.Error == nil {
		log.ID = int64(orderModel.ID)
	}
	return res.Error
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(postgres.Open(dataSourceUrl), &gorm.Config{})
	// db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("log"))); err != nil {
		return nil, fmt.Errorf("db otel plugin error: %v", err)
	}

	err := db.AutoMigrate(&Log{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}
