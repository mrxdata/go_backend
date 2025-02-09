package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func SessionProtected(c fiber.Ctx) error {
	sess := session.FromContext(c)
	if sess == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	isAuthorized := sess.Get("authorized")
	fmt.Println("sess data")
	fmt.Println(sess.ID())
	fmt.Println(sess.Get("uuid"))
	fmt.Println(sess.Get("role"))
	fmt.Println(sess.Get("authorized"))
	fmt.Println(sess.Get("signed_via"))
	fmt.Printf("isAuthorized: %s", isAuthorized)
	if isAuthorized != "true" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "User not authenticated!",
		})
	}
	sess.Session.Save()
	return c.Next()
}

func SessionMiddleware(app *fiber.App, store *session.Store, middleware fiber.Handler) {
	app.Use(
		middleware,
		//nolint // ignore error
		//csrf.New(
		//	csrf.Config{ // TODO: config
		//		Session: store,
		//		ErrorHandler: func(c fiber.Ctx, err error) error {
		//			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		//				"error": "CSRF token is missing or invalid",
		//			})
		//		},
		//	},
		//),
	)

}
