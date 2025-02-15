package main

import (
	"context"
	"fmt"

	"github.com/G-Villarinho/fast-feet-api/config"
	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/handlers"
	"github.com/G-Villarinho/fast-feet-api/middlewares"
	"github.com/G-Villarinho/fast-feet-api/repositories"
	"github.com/G-Villarinho/fast-feet-api/services"
	"github.com/G-Villarinho/fast-feet-api/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()
	config.ConfigureLogger()

	ctx := context.Background()

	e := echo.New()
	i := di.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	middlewares.Cors(e)

	DB, err := storage.NewPostgresStorage(ctx)
	if err != nil {
		e.Logger.Fatal(err)
	}

	di.Provide(i, func(d *di.Injector) (*gorm.DB, error) {
		return DB, nil
	})

	di.Provide(i, handlers.NewAuthHandler)
	di.Provide(i, handlers.NewOrderHandler)
	di.Provide(i, handlers.NewRecipientHandler)
	di.Provide(i, handlers.NewUserHandler)

	di.Provide(i, services.NewAuthService)
	di.Provide(i, services.NewFileService)
	di.Provide(i, services.NewOrderService)
	di.Provide(i, services.NewRecipientService)
	di.Provide(i, services.NewSecureService)
	di.Provide(i, services.NewTokenService)
	di.Provide(i, services.NewUserService)

	di.Provide(i, repositories.NewOrderRepository)
	di.Provide(i, repositories.NewRecipientRepository)
	di.Provide(i, repositories.NewUserRepository)

	if err := handlers.SetupRoutes(e, i); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Env.API.Port)))
}
