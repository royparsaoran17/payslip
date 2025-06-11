package order

import (
	"context"
	"payroll-se/internal/common"
	"payroll-se/internal/presentations"

	"payroll-se/internal/entity"
)

type Order interface {
	GetAllOrder(ctx context.Context, userID string, meta *common.Metadata) ([]entity.Order, error)
	GetOrderByID(ctx context.Context, orderID string) (*entity.OrderDetail, error)
	CreateOrder(ctx context.Context, input presentations.Order) (*entity.Order, error)
	CreateOrderPayment(ctx context.Context, input presentations.OrderPayment) (*entity.Order, error)
}
