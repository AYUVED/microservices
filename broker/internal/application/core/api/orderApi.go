package orderApi

import (
	"context"
	"github.com/ayuved/microservices-helper/domain"

	"github.com/ayuved/microservices/broker/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	order ports.OrderPort
}

func NewApplication(order ports.OrderPort) *Application {
	return &Application{
		order: order,
	}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {

	orderErr := a.order.PlaceOrder(ctx, &order)
	if orderErr != nil {
		st, _ := status.FromError(orderErr)
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "order",
			Description: st.Message(),
		}
		// var allErrors []string
		// for _, detail := range st.Details() {
		// 	switch t := detail.(type) {
		// 	case *errdetails.BadRequest:
		// 		for _, violation := range t.GetFieldViolations() {
		// 			allErrors = append(allErrors, violation.Description)
		// 		}
		// 	}
		// }

		// fieldErr := &errdetails.BadRequest_FieldViolation{
		// 	Field:       "payment",
		// 	Description: strings.Join(allErrors, "\n"),
		// }
		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusWithDetails, _ := orderStatus.WithDetails(badReq)
		return domain.Order{}, statusWithDetails.Err()
	}
	return order, nil
}

func (a Application) GetOrder(ctx context.Context, id int64) (domain.Order, error) {
	return a.order.GetOrder(ctx, id)
}
