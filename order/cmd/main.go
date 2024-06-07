package main

import (
	"log"

	"GitHub.com/ayuved/microservices/order/internal/application/core/api"
	"Github.com/ayuved/microservices/order/config"
	"Github.com/ayuved/microservices/order/internal/adapters/db"
	"github.com/ayuved/microservices/order/internal/adapters/grpc"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
