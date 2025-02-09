package controllers

import (
	"flutty_messenger/app/models"
	"flutty_messenger/pkg/repository"
	"flutty_messenger/pkg/utils"
	"flutty_messenger/platform/database"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func TgAuthHandler(c fiber.Ctx) error {
	tgAuthData := models.TgAuthData{}

	if err := c.Bind().Body(&tgAuthData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validateTg := utils.NewTgValidator()

	if err := validateTg.Struct(tgAuthData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	postgresdb, err := database.OpenDBConnection(database.PostgreSQL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	user, err := postgresdb.GetUserByTgUserID(tgAuthData.Id)
	if err != nil {
		log.Errorf("Can't get user by id: %v", err)
		user.SignedVia = models.SignedViaTg
		user.Role = repository.RoleUser
		user.UpdateUserTgInfo(tgAuthData)

		user.ID, err = postgresdb.CreateUser(&user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	}

	sess := session.FromContext(c)
	if sess == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "session is nil",
		})
	}

	sess.Set("uuid", user.ID.String())
	sess.Set("role", user.Role)
	sess.Set("authorized", true)
	sess.Set("signed_via", user.SignedVia)
	fmt.Println("sess data")
	fmt.Println(sess.ID())
	fmt.Println(sess.Get("uuid"))
	fmt.Println(sess.Get("role"))
	fmt.Println(sess.Get("authorized"))
	fmt.Println(sess.Get("signed_via"))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Authorized successfully",
	})
}
