package services

import (
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"
)

type orderService struct {
	orderRepo ports.OrderRepository
}

func NewOrderService(orderRepo ports.OrderRepository) ports.OrderService {
	return orderService{orderRepo}
}

func (s orderService) Create(order *model_gorm.Order) error {
	err := s.orderRepo.Create(order)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s orderService) UpdateStatus(orderId *int, status *string) error {
	err := s.orderRepo.UpdateStatus(orderId, status)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s orderService) GetPendingOrders(userId *int) (*[]model_io.Order, error) {
	orders, err := s.orderRepo.GetPendingOrders(userId)
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	return orders, nil
}

func (s orderService) GetOrders(userId *int) (*[]model_io.Order, error) {
	orders, err := s.orderRepo.GetOrders(userId)
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	return orders, nil
}

func (s orderService) GetOrderById(orderId *int) (*model_io.Order, error) {
	order, err := s.orderRepo.GetOrderById(orderId)
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	return order, nil
}
