package router

import (
	"github.com/gofiber/fiber/v2"
	"ozanpay/database"
	"ozanpay/handler"
	"ozanpay/service"
)

func Router(app *fiber.App) {
	db := database.DB()
	userService := service.NewUserService(db)

	userHandler := handler.NewUserHandler(userService)

	app.Post("/users", userHandler.Create)
}
