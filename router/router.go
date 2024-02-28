package router

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"ozanpay/config"
	"ozanpay/database"
	"ozanpay/handler"
	"ozanpay/middleware"
	"ozanpay/service"
)

func Router(app *fiber.App, cfg config.ServerConfig) {
	db := database.DB()
	userService := service.NewUserService(db)

	userHandler := handler.NewUserHandler(userService)

	api := app.Group("/api")
	api.Use(jwtware.New(jwtware.Config{
		SigningKey:     []byte(cfg.JwtSecret),
		ErrorHandler:   middleware.JwtErrorHandler,
		SuccessHandler: middleware.JwtSuccessHandler,
		TokenLookup:    "header:" + fiber.HeaderAuthorization + ",query:token",
	}))

	app.Post("/users", userHandler.Create)
}
