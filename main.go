package main

import (
	"context"
	"fmt"
	"go-ecommerce-demo/adapters/handlers"
	"go-ecommerce-demo/adapters/repositories"
	"go-ecommerce-demo/core/services"
	"go-ecommerce-demo/middlewares"
	"go-ecommerce-demo/routes"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
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

	app.Use(authCards, middlewares.JwtWare(userSrv))

	app.Use(adminCards, middlewares.Role)

	routes.InitRoutes(app, userHandle, productHandle, catHandle, cartHandle, orderHandle)

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
