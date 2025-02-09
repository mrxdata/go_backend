package routes

import (
	"flutty_messenger/app/controllers"
	"github.com/gofiber/fiber/v3"
)

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api")

	route.Post("/auth/telegram", controllers.TgAuthHandler)
	route.Post("/info", controllers.BeaconInfoHandler)
	route.Post("/test", controllers.TestHandler)

	route.Get("/test", controllers.TestHandler)
	route.Get("/test/db", controllers.HealthCheckDB)
}
