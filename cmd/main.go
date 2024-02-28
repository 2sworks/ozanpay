package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"ozanpay/config"
	"ozanpay/database"
	"ozanpay/router"
	"syscall"
)

func main() {
	cfg, err := config.Setup()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Use(cors.New())

	app.Use(healthcheck.New())
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

	database.ConnectAndMigrate(cfg.Database)
	router.Router(app, cfg.Server)
	go func() {
		log.Info("Bismillah")
		if errr := app.Listen(fmt.Sprintf(":%v", cfg.Server.Port)); errr != nil {
			log.Error("fiber patladı", zap.Error(errr))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c // block until interr
	log.Info("Gracefully shutting down")
	_ = app.Shutdown()
	log.Info("Elhamdülillah")
}
