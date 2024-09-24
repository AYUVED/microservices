package grpc

import (
	"context"

	"github.com/ayuved/microservices-proto/golang/order"
	"github.com/ayuved/microservices-helper/domain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	log.WithContext(ctx).Info("Creating order...")
	var validationErrors []*errdetails.BadRequest_FieldViolation
	if request.UserId == 0 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "user_id",
			Description: "user_id is required",
		})
	}
	if len(request.OrderItems) == 0 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "order_items",
			Description: "order_items is required",
		})
	}
	if len(validationErrors) > 0 {
		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = validationErrors
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusWithDetails, _ := orderStatus.WithDetails(badReq)
		return nil, statusWithDetails.Err()
	}

	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	log.WithContext(ctx).Infof("OrderItems: %v\n", orderItems)
	newOrder := domain.NewOrder(request.UserId, orderItems)
	log.WithContext(ctx).Infof("NewOrder: %v\n", newOrder)
	result, err := a.api.PlaceOrder(ctx, newOrder)
	log.WithContext(ctx).Infof("Result: %v\n", result)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}

func (a Adapter) Get(ctx context.Context, request *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	result, err := a.api.GetOrder(ctx, request.OrderId)
	var orderItems []*order.OrderItem
	for _, orderItem := range result.OrderItems {
		orderItems = append(orderItems, &order.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	if err != nil {
		return nil, err
	}
	return &order.GetOrderResponse{UserId: result.CustomerID, OrderItems: orderItems}, nil
}
