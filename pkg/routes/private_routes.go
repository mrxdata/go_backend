package routes

import (
	"flutty_messenger/app/controllers"
	"flutty_messenger/pkg/middleware"
	"github.com/gofiber/fiber/v3"
)

func PrivateRoutes(app *fiber.App) {
	route := app.Group("/api")

	route.Post("/testpr", controllers.TestHandler, middleware.SessionProtected)
	route.Get("/testpr", controllers.TestHandler, middleware.SessionProtected)
}
