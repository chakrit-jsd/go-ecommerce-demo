package model_gorm

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique;size:20;not null"`
	Password string `gorm:"size:60;not null"`
	Address  string `gorm:"size:255;not null"`
	Role     string `gorm:"size:5;default:user"`

	CartDetails []CartDetail `gorm:"foreignKey:UserId;refernces:ID"`
	Order       []Order
}

type Category struct {
	gorm.Model
	Name string `gorm:"unique;size:40;not null"`
}

type Product struct {
	gorm.Model
	Name       string  `gorm:"unique;size:20;not null"`
	Detail     string  `gorm:"size:100;not null"`
	Stock      uint    `gorm:"not null"`
	Price      float64 `gorm:"type:decimal(16,2);not null"`
	CategoryId int
	Category   Category
}

// type Cart struct {
// 	gorm.Model
// 	UserId      int
// 	CartDetails []CartDetail
// }

type CartDetail struct {
	UserId    int
	ProductId int
	Product   Product
	Quantity  int `gorm:"not null"`
}

type Order struct {
	gorm.Model
	UserId       int           `json:"user_id"`
	OrderDetails []OrderDetail `json:"order_details"`
	Status       string        `gorm:"type:enum('pending', 'cancelled', 'confirmed');default:pending"`
}

type OrderDetail struct {
	OrderId   int
	ProductId int `json:"product_id"`
	Product   Product
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"type:decimal(16,2);not null" json:"price"`
}

type OrderStatusEnum struct {
	Cancel, Confirm, Pending string
}

var Enum = OrderStatusEnum{
	Cancel:  "canceled",
	Confirm: "confirmed",
	Pending: "pending",
}
