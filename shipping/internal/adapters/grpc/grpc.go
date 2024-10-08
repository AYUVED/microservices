package grpc

import (
	"context"
	"fmt"

	"github.com/ayuved/microservices-proto/golang/shipping"
	"github.com/ayuved/microservices-helper/domain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	log.WithContext(ctx).Info("Creating shipping...")
	var shipItems []domain.ShippingItem
	for _, orderItem := range request.ShippingItems {
		shipItems = append(shipItems, domain.ShippingItem{

			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newShipping := domain.NewShipping(request.OrderId, request.AddressId, shipItems)
	result, err := a.api.Charge(ctx, newShipping)
	log.WithContext(ctx).Infof("Shipping result: %v\n", result)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &shipping.CreateShippingResponse{ShippingId: result.ID}, nil
}
