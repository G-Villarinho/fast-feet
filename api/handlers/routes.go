package handlers

import (
	"fmt"

	"github.com/G-Villarinho/fast-feet-api/di"
	"github.com/G-Villarinho/fast-feet-api/middlewares"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, i *di.Injector) error {
	if err := SetupUserRoutes(e, i); err != nil {
		return fmt.Errorf("setup user routes: %w", err)
	}

	if err := SetupAuthRoutes(e, i); err != nil {
		return fmt.Errorf("setup auth routes: %w", err)
	}

	if err := SetupRecipientRoutes(e, i); err != nil {
		return fmt.Errorf("setup recipient routes: %w", err)
	}

	if err := SetupOrdersRoutes(e, i); err != nil {
		return fmt.Errorf("setup order routes: %w", err)
	}

	return nil
}

func SetupUserRoutes(e *echo.Echo, i *di.Injector) error {
	h, err := di.Invoke[UserHandler](i)
	if err != nil {
		return fmt.Errorf("invoke user handler: %w", err)
	}

	v1Group := e.Group("/v1/users", middlewares.Authenticate)

	v1Group.POST("/admin", h.CreateAdmin)
	v1Group.GET("/me", h.GetUser)

	return nil
}

func SetupAuthRoutes(e *echo.Echo, i *di.Injector) error {
	h, err := di.Invoke[AuthHandler](i)
	if err != nil {
		return fmt.Errorf("invoke auth handler: %w", err)
	}

	v1Group := e.Group("/v1")

	v1Group.POST("/login", h.Login)
	v1Group.POST("/logout", h.Logout)

	return nil
}

func SetupRecipientRoutes(e *echo.Echo, i *di.Injector) error {
	h, err := di.Invoke[RecipientHandler](i)
	if err != nil {
		return fmt.Errorf("invoke recipient handler: %w", err)
	}

	v1Group := e.Group("/v1/recipients", middlewares.Authenticate)

	v1Group.POST("", h.CreateRecipient)
	v1Group.GET("/:recipientId", h.GetRecipient)
	v1Group.GET("/lite", h.GetRecipientsBasicInfo)
	v1Group.DELETE("/:recipientId", h.DeleteRecipient)
	v1Group.PUT("/:recipientId", h.UpdateRecipient)

	return nil
}

func SetupOrdersRoutes(e *echo.Echo, i *di.Injector) error {
	h, err := di.Invoke[OrderHandler](i)
	if err != nil {
		return fmt.Errorf("invoke order handler: %w", err)
	}

	v1Group := e.Group("/v1/orders", middlewares.Authenticate)

	v1Group.POST("", h.CreateOrder)
	v1Group.PATCH("/:orderId/status/pick-up", h.PickUpOrder)
	v1Group.PATCH("/:orderId/status/deliver", h.DeliverOrder)
	v1Group.GET("", h.GetOrders)

	return nil
}
