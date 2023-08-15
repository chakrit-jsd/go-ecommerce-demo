package repositories

import (
	"errors"
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"

	"gorm.io/gorm"
)

type orderRepositoryDB struct {
	db *gorm.DB
}

func NewOrderRepositoryDB(db *gorm.DB) ports.OrderRepository {
	return orderRepositoryDB{db}
}

func (r orderRepositoryDB) Create(order *model_gorm.Order) error {
	err := r.db.Model(&model_gorm.Order{}).Create(order).Error
	if err != nil {
		return err
	}

	return nil
}

func (r orderRepositoryDB) UpdateStatus(orderId *int, status *string) error {
	res := r.db.Model(&model_gorm.Order{}).Where("id = ?", orderId).Update("status", status)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (r orderRepositoryDB) GetPendingOrders(userId *int) (*[]model_io.Order, error) {
	orders := []model_io.Order{}
	err := r.db.Model(&model_gorm.Order{}).
		Where("user_id = ?", *userId).
		Where("status = ?", model_gorm.Enum.Pending).
		Preload("OrderDetails.Product.Category").
		Order("created_at DESC").
		Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (r orderRepositoryDB) GetOrders(userId *int) (*[]model_io.Order, error) {
	orders := []model_io.Order{}
	err := r.db.Model(&model_gorm.Order{}).
		Where("user_id = ?", *userId).
		Where("status != ?", model_gorm.Enum.Pending).
		Preload("OrderDetails.Product.Category").
		Order("created_at DESC").
		Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (r orderRepositoryDB) GetOrderById(orderId *int) (*model_io.Order, error) {
	order := model_io.Order{}
	err := r.db.Model(&model_gorm.Order{}).
		Where("id = ?", orderId).
		Preload("OrderDetails.Product.Category").
		First(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}
