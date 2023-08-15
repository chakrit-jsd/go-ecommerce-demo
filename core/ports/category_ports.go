package ports

import (
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
)

type CategoryRepository interface {
	Create(*model_gorm.Category) error
	Update(*model_gorm.Category, *int) error
	Delete(*int) error
	GetCategories() (*[]model_io.Category, error)
	GetCategoryById(*int) (*model_io.Category, error)
	// GetCategoryByName(*string) (*model_io.Category, error)
}

type CategoryService interface {
	Create(*model_io.Category) error
	Update(*model_io.Category, *int) error
	Delete(*int) error
	GetCategories() (*[]model_io.Category, error)
	GetCategoryById(*int) (*model_io.Category, error)
}
