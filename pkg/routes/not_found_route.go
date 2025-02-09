package routes

import "github.com/gofiber/fiber/v3"

func NotFoundRoute(app *fiber.App) {
	app.Use(
		func(c fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "Invalid endpoint",
			})
		},
	)
}
