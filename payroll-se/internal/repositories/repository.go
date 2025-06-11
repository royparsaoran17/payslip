package repositories

import (
	"context"
	"database/sql"
	"payroll-se/internal/repositories/internal/order"
	"payroll-se/internal/repositories/internal/orderitem"
	"payroll-se/internal/repositories/internal/payment"
	"payroll-se/internal/repositories/internal/product"
	"payroll-se/internal/repositories/internal/stockreservation"
	"payroll-se/internal/repositories/internal/user"
	"payroll-se/internal/repositories/internal/warehouse"
	"payroll-se/pkg/databasex"
)

type Repository struct {
	Order            Order
	OrderItem        OrderItem
	Payment          Payment
	StockReservation StockReservation
	User             User
	Warehouse        Warehouse
	Product          Product
	db               databasex.Adapter
}

func NewRepository(db databasex.Adapter) *Repository {
	return &Repository{
		Order:            order.NewOrder(db),
		Payment:          payment.NewPayment(db),
		OrderItem:        orderitem.NewOrderItem(db),
		StockReservation: stockreservation.NewStockReservation(db),
		User:             user.NewUser(db),
		Product:          product.NewProduct(db),
		Warehouse:        warehouse.NewWarehouse(db),
	}
}

func (r Repository) BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, options)
}
