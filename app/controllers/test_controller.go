package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
)

func TestHandler(c fiber.Ctx) error {
	fmt.Println("TestHandler")

	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}
