package ports

import (
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
)

type CartRepository interface {
	GetCart(*int) (*[]model_io.CartDetail, error)
	CreateCart(*int) error
	AddProduct(*model_gorm.CartDetail) error
	UpdateProduct(*model_gorm.CartDetail) error
	DeleteProduct(*model_io.DeleteProductsInCart) error
}

type CartService interface {
	GetCart(*int) (*[]model_io.CartDetail, error)
	AddProduct(*model_io.CartDetail) error
	UpdateProduct(*model_io.CartDetail) error
	DeleteProduct(*model_io.DeleteProductsInCart) error
}
