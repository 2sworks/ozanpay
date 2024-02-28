package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"ozanpay/database"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(healthcheck.New())

	// Or extend your config for customization
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		ReadinessEndpoint: "/ready",
	}))

	database.ConnectAndMigrate()
	app.Listen(fmt.Sprintf(":%s", "3000"))
}
