package configs

import (
	"github.com/bytedance/sonic"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

func FiberConfig() fiber.Config {
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	}
}
