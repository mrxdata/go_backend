package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"os"
)

func FiberMiddleware(app *fiber.App) {
	// TO DO
	// add compress and config zstd and minsize (also do it in nginx.conf)
	// add recover.new
	app.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://proxy"},
			AllowMethods:     []string{"GET,POST,OPTIONS"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: true,
		}),
		logger.New(logger.Config{
			Format:     "${time} | ${pid} | ${status} | ${latency} | ${ip} | ${ips} | ${host} | ${protocol} | ${port} | ${method} | ${url} | ${queryParams} | Sent: ${bytesSent} | Recv: ${bytesReceived} | ${error}\n",
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Local",
			Output:     os.Stdout,
		}),
	)
}
