package routes

import (
	"go-ecommerce-demo/adapters/handlers"

	"github.com/gofiber/fiber/v2"
)

type routers struct {
	auth,
	users,
	products,
	categories,
	carts,
	orders fiber.Router
}

type handles struct {
	userHandle    handlers.UserHandler
	productHandle handlers.ProductHandler
	catHandle     handlers.CategoryHandler
	cartHandle    handlers.CartHandler
	orderHandle   handlers.OrderHandler
}

func InitRoutes(app *fiber.App, user handlers.UserHandler, prod handlers.ProductHandler, cate handlers.CategoryHandler, cart handlers.CartHandler, order handlers.OrderHandler) {
	routers := routers{
		auth:       app.Group("/auth"),
		users:      app.Group("/users"),
		products:   app.Group("/products"),
		categories: app.Group("/categories"),
		carts:      app.Group("/carts"),
		orders:     app.Group("/orders"),
	}

	handles := handles{
		userHandle:    user,
		productHandle: prod,
		catHandle:     cate,
		cartHandle:    cart,
		orderHandle:   order,
	}

	initSubRoutes(routers, handles)
}
