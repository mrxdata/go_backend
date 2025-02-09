package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v3"
)

func StartServerWithGracefulShutdown(a *fiber.App, config fiber.ListenConfig) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Can't shutdown server. Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	if err := a.Listen(fiberConnURL, config); err != nil {
		log.Printf("Can't start server. Reason: %v", err)
	}

	<-idleConnsClosed
}

func StartServer(a *fiber.App, config fiber.ListenConfig) {
	fiberConnURL, _ := ConnectionURLBuilder("fiber")
	if err := a.Listen(fiberConnURL, config); err != nil {
		log.Printf("Can't start server. Reason: %v", err)
	}
}
