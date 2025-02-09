package main

import (
	"github.com/gofiber/fiber/v3/middleware/session"
	"os"

	"flutty_messenger/pkg/configs"
	"flutty_messenger/pkg/middleware"
	"flutty_messenger/pkg/routes"
	"flutty_messenger/pkg/utils"

	"github.com/gofiber/fiber/v3"

	_ "flutty_messenger/docs"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fiberConfig := configs.FiberConfig()
	listenConfig := configs.ListenConfig()
	sessionConfig := configs.SessionConfig()

	app := fiber.New(fiberConfig)

	sessionMiddleware, SessionStore := session.NewWithStore(sessionConfig) // TODO: config

	middleware.FiberMiddleware(app)
	middleware.SessionMiddleware(app, SessionStore, sessionMiddleware)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)
	routes.SwaggerRoute(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app, listenConfig)
	} else {
		utils.StartServerWithGracefulShutdown(app, listenConfig)
	}
}
