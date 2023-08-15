package ports

import (
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
)

type OrderRepository interface {
	Create(*model_gorm.Order) error
	UpdateStatus(orderId *int, status *string) error
	GetPendingOrders(userId *int) (*[]model_io.Order, error)
	GetOrders(userId *int) (*[]model_io.Order, error)
	GetOrderById(orderId *int) (*model_io.Order, error)
}

type OrderService interface {
	Create(*model_gorm.Order) error
	UpdateStatus(orderId *int, status *string) error
	GetPendingOrders(userId *int) (*[]model_io.Order, error)
	GetOrders(userId *int) (*[]model_io.Order, error)
	GetOrderById(orderId *int) (*model_io.Order, error)
}
