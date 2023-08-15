package model_io

type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Detail     string   `json:"detail"`
	Stock      uint     `json:"stock"`
	Price      float64  `json:"price"`
	CategoryId int      `json:"category_id"`
	Category   Category `json:"category"`
}

// type Cart struct {
// 	ID          int          `json:"id"`
// 	UserId      int          `json:"owner"`
// 	CartDetails []CartDetail `json:"detail"`
// }

type CartDetail struct {
	UserId    int     `json:"user_id"`
	ProductId int     `json:"product_id"`
	Product   Product `json:"products"`
	Quantity  int     `json:"quantity"`
	// Price     float64 `json:"price"`
}

type DeleteProductsInCart struct {
	UserId     int   `json:"user_id"`
	ProductsId []int `json:"products_id"`
}

type Order struct {
	ID           int           `json:"id"`
	UserId       int           `json:"owner"`
	CreatedAt    string        `json:"created_at"`
	OrderDetails []OrderDetail `json:"detail"`
	Status       string        `json:"status"`
}

type OrderDetail struct {
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Product   Product `json:"products"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type ProductsAndCounts struct {
	Products    []Product `json:"products"`
	TotalCounts int64     `json:"counts"`
}

type QueryProducts struct {
	Name       string `params:"name"`
	CategoryId int    `params:"catid"`
	Sort       string `params:"sort"`
	Page       int    `params:"page"`
	Offset     int    `params:"offset"`
	Count      bool   `params:"count"`
}
