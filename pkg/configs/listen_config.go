package configs

import "github.com/gofiber/fiber/v3"

func ListenConfig() fiber.ListenConfig {

	return fiber.ListenConfig{
		EnablePrefork: true,
	}
}
