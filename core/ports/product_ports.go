package ports

import (
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
)

type ProductRepository interface {
	Create(*model_gorm.Product) error
	Update(*model_gorm.Product, *int) error
	UpdateStock(quantity, productId *int) error
	Delete(*int) error
	GetProducts() (*[]model_io.Product, error)
	GetProductById(*int) (*model_io.Product, error)
	GetProductsByQuery(*model_io.QueryProducts) (*[]model_io.Product, error)
	GetCounts(*model_io.QueryProducts) (*int64, error)
}

type ProductService interface {
	Create(*model_io.Product) error
	Update(*model_io.Product, *int) error
	UpdateStock(quantity, productId *int) error
	Delete(*int) error
	GetProducts() (*[]model_io.Product, error)
	GetProductById(*int) (*model_io.Product, error)
	GetProductsByQuery(*model_io.QueryProducts, bool) (*model_io.ProductsAndCounts, error)
}
