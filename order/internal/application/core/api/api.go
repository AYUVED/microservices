package api

import (
	"context"
	"log"

	"github.com/ayuved/microservices/order/internal/application/core/domain"
	"github.com/ayuved/microservices/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	log.Printf("PlaceOrder1: %v\n", order)
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	log.Printf("PlaceOrder2: %v\n", order)
	paymentErr := a.payment.Charge(ctx, &order)
	if paymentErr != nil {
		st, _ := status.FromError(paymentErr)
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
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
	log.Printf("Order: %v\n", order)
	return order, nil
}

func (a Application) GetOrder(ctx context.Context, id int64) (domain.Order, error) {
	return a.db.Get(ctx, id)
}
