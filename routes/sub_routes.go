package routes

func initSubRoutes(r routers, h handles) {
	r.auth.Post("/register", h.userHandle.PostRegister)
	r.auth.Post("/login", h.userHandle.PostLogin)

	r.users.Get("/me", h.userHandle.GetMe)
	r.users.Put("/update", h.userHandle.PutUpdate)
	r.users.Delete("/delete", h.userHandle.DeleteUser)

	r.products.Get("/:name<string>?-:catid<uint>?-:sort<string>?-:page<uint>?-:offset<uint>?-:count<bool>?", h.productHandle.GetProducts)
	r.products.Post("/create", h.productHandle.PostCreateProduct)
	r.products.Put("/update", h.productHandle.PutUpdateProduct)
	r.products.Delete("/delete", h.productHandle.DeleteProduct)

	r.categories.Get("/all", h.catHandle.GetCategories)
	r.categories.Post("/create", h.catHandle.PostCreateCategory)
	r.categories.Put("/update", h.catHandle.PutUpdateCategory)
	r.categories.Delete("/delete", h.catHandle.DeleteCategory)

	r.carts.Get("/me", h.cartHandle.GetCart)
	r.carts.Post("/add", h.cartHandle.PostAddProduct)
	r.carts.Put("/update", h.cartHandle.PutUpdateProduct)
	r.carts.Delete("/delete", h.cartHandle.DeleteProduct)

	r.orders.Post("/create", h.orderHandle.PostCreateOrder)
	r.orders.Get("/pending", h.orderHandle.GetPendingOrders)
	r.orders.Put("/update", h.orderHandle.PutUpdateStatus)
	r.orders.Get("/history", h.orderHandle.GetHistory)
}
