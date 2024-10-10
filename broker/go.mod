module github.com/ayuved/microservices/broker

go 1.22.3

require (
	github.com/ayuved/microservices-helper v0.0.0-20241010022403-ac39f80bde3e
	github.com/ayuved/microservices-proto/golang/order v1.0.9
	github.com/go-chi/chi/v5 v5.0.14
	github.com/go-chi/cors v1.2.1
	github.com/rabbitmq/amqp091-go v1.10.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.55.0
	google.golang.org/grpc v1.66.2
)

require (
	github.com/ayuved/microservices-proto/golang/logservice v1.0.9 // indirect
	github.com/ayuved/microservices-proto/golang/shipping v1.0.9 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/sirupsen/logrus v1.9.3
	go.opentelemetry.io/otel v1.30.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1
	google.golang.org/protobuf v1.34.2 // indirect
)
