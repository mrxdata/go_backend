package controllers

import (
	"flutty_messenger/platform/database"
	"github.com/gofiber/fiber/v3"
)

func HealthCheckDB(c fiber.Ctx) error {
	postgresdb, err := database.OpenDBConnection(database.PostgreSQL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to connect to database: " + err.Error(),
		})
	}
	defer func(postgresdb *database.Queries) {
		err := postgresdb.Close()
		if err != nil {

		}
	}(postgresdb)

	err = postgresdb.Ping()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Database ping failed: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Database connection is healthy!",
	})
}
