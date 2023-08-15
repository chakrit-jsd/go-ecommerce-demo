package main

import (
	"context"
	"fmt"
	"go-ecommerce-demo/adapters/handlers"
	"go-ecommerce-demo/adapters/repositories"
	"go-ecommerce-demo/core/services"
	"go-ecommerce-demo/middlewares"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()

	userRepo := repositories.NewUserRepositoryDB(db)
	userSrv := services.NewUserService(userRepo)
	userHandle := handlers.NewUserHandler(userSrv)

	productRepo := repositories.NewProductRepositoryDB(db)
	productSrv := services.NewProductService(productRepo)
	productHandle := handlers.NewProductHandle(productSrv)

	categoryRepo := repositories.NewCategoryRepositoryDB(db)
	categorySrv := services.NewCategoryService(categoryRepo)
	catHandle := handlers.NewCategoryHandle(categorySrv)

	cartRepo := repositories.NewCartRepositoryDB(db)
	cartSrv := services.NewCartService(cartRepo)
	cartHandle := handlers.NewCartHandle(cartSrv)

	orderRepo := repositories.NewOrderRepositoryDB(db)
	orderSrv := services.NewOrderService(orderRepo)
	orderHandle := handlers.NewOrderHandle(orderSrv, cartSrv, productSrv)

	app := fiber.New(fiber.Config{
		Prefork:        false,
		RequestMethods: fiber.DefaultMethods,
	})
	authCards := []string{
		"/users",
		"/products/create",
		"/products/update",
		"/products/delete",
		"/categories/create",
		"/categories/update",
		"/categories/delete",
		"/carts",
		"/orders",
	}

	adminCards := authCards[1:7]

	app.Use(authCards, jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("jwt.secret"))},
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)
			str, err := token.Claims.GetIssuer()
			if err != nil {
				return fiber.ErrBadRequest
			}

			userId, err := strconv.Atoi(str)
			if err != nil {
				return fiber.ErrInternalServerError
			}

			userInfo, err := userSrv.GetUserById(&userId)
			if err != nil {
				return fiber.ErrUnauthorized
			}
			c.Locals("userInfo", userInfo)
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.ErrUnauthorized
		},
	}))

	app.Use(adminCards, middlewares.Role)

	auth := app.Group("/auth")
	users := app.Group("/users")
	products := app.Group("/products")
	categories := app.Group("/categories")
	carts := app.Group("/carts")
	orders := app.Group("/orders")
	// db.Migrator().AutoMigrate(&model_gorm.User{})
	auth.Post("/register", userHandle.PostRegister)
	auth.Post("/login", userHandle.PostLogin)

	users.Get("/me", userHandle.GetMe)
	users.Put("/update", userHandle.PutUpdate)
	users.Delete("/delete", userHandle.DeleteUser)

	products.Get("/:name<string>?-:catid<uint>?-:sort<string>?-:page<uint>?-:offset<uint>?-:count<bool>?", productHandle.GetProducts)
	products.Post("/create", productHandle.PostCreateProduct)
	products.Put("/update", productHandle.PutUpdateProduct)
	products.Delete("/delete", productHandle.DeleteProduct)

	categories.Get("/all", catHandle.GetCategories)
	categories.Post("/create", catHandle.PostCreateCategory)
	categories.Put("/update", catHandle.PutUpdateCategory)
	categories.Delete("/delete", catHandle.DeleteCategory)

	carts.Get("/me", cartHandle.GetCart)
	carts.Post("/add", cartHandle.PostAddProduct)
	carts.Put("/update", cartHandle.PutUpdateProduct)
	carts.Delete("/delete", cartHandle.DeleteProduct)

	orders.Post("/create", orderHandle.PostCreateOrder)
	orders.Get("/pending", orderHandle.GetPendingOrders)
	orders.Put("/update", orderHandle.PutUpdateStatus)
	orders.Get("/history", orderHandle.GetHistory)

	app.Listen(fmt.Sprintf("%v:%v",
		viper.GetString("app.host"),
		viper.GetString("app.port"),
	))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n*****************************\n", sql)
}

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		// viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)
	dial := mysql.Open(dsn)

	db, err := gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	return db
}
