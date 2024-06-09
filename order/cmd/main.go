package main

import (
	"log"

	"github.com/ayuved/microservices/order/config"
	"github.com/ayuved/microservices/order/internal/adapters/db"
	"github.com/ayuved/microservices/order/internal/adapters/grpc"
	"github.com/ayuved/microservices/order/internal/application/core/api"
)

func main() {
	log.Println("Starting order service...")
	log.Println("Connecting to database..." + config.GetDataSourceURL())
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
