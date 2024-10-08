package db

import (
	"context"
	"fmt"
	"log"
	"testing" // Add missing import for "testing"

_ "github.com/lib/pq"
	"github.com/ayuved/microservices-helper/domain"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type OrderDatabaseTestSuite struct {
	suite.Suite
	DataSourceUrl string
}

func (o *OrderDatabaseTestSuite) SetupSuite() {
	ctx := context.Background()
	port := "5432/tcp"
	dbURL := func(port nat.Port) string {
		return fmt.Sprintf("postgres://postgres:password@localhost:%s/orders?sslmode=disable", port.Port())
	}
	req := testcontainers.ContainerRequest{
		Image:        "docker.io/postgres:15.0",
		ExposedPorts: []string{port},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":      "orders",
		},
		WaitingFor: wait.ForSQL(nat.Port(port), "postgres", func(host string, port nat.Port) string {
			return dbURL(port)
		}),
	}
	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal("Failed to start Mysql.", err)
	}
	endpoint, _ := postgresContainer.Endpoint(ctx,"")
	o.DataSourceUrl = fmt.Sprintf("postgres://postgres:password@%s/orders?sslmode=disable", endpoint)
}

func (o *OrderDatabaseTestSuite) Test_Should_Save_Order() {
	adapter, err := NewAdapter(o.DataSourceUrl)
	o.Nil(err)
	saveErr := adapter.Save(context.Background(), &domain.Order{})
	o.Nil(saveErr)
}

func (o *OrderDatabaseTestSuite) Test_Should_Get_Order() {
	adapter, _ := NewAdapter(o.DataSourceUrl)
	order := domain.NewOrder(2, []domain.OrderItem{
		{
			ProductCode: "CAM",
			Quantity:    5,
			UnitPrice:   1.32,
		},
	})
	ctx := context.Background()
	adapter.Save(ctx, &order)
	ord, _ := adapter.Get(ctx, order.ID)
	o.Equal(int64(2), ord.CustomerID)
}

func TestOrderDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDatabaseTestSuite))
}
