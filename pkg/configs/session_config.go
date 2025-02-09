package configs

import (
	"github.com/gofiber/fiber/v3/middleware/session"
)

func SessionConfig() session.Config {

	return session.Config{
		CookieHTTPOnly: true,
		CookieSameSite: "Strict",
		CookieSecure:   false,
	}
}
