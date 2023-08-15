package ports

import (
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
)

type UserRepository interface {
	Create(*model_gorm.User) error
	Update(*model_gorm.User, *int) error
	Delete(*int) error
	GetUserById(*int) (*model_io.User, error)
	GetUserByUserName(*string) (*model_io.User, error)
}

type UserService interface {
	Create(*model_io.User) error
	Update(*model_io.User, *int) error
	Delete(*int) error
	GetUserById(*int) (*model_io.User, error)
	GetUserByUserName(*string) (*model_io.User, error)
	Login(*model_io.User) (*string, error)
}
