package domain

import (
	"time"
)

type Log struct {
	ID         int64   `json:"id"`
	CustomerID int64   `json:"customer_id"`
	Status     string  `json:"status"`
	OrderId    int64   `json:"order_id"`
	TotalPrice float32 `json:"total_price"`
	CreatedAt  int64   `json:"created_at"`
}

func NewLog(customerId int64, orderId int64, totalPrice float32) Log {
	return Log{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CustomerID: customerId,
		OrderId:    orderId,
		TotalPrice: totalPrice,
	}
}
