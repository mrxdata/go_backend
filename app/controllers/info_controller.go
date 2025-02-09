package controllers

import (
	"encoding/json"
	"flutty_messenger/pkg/utils"
	"github.com/gofiber/fiber/v3"
	"log"
)

func BeaconInfoHandler(c fiber.Ctx) error {
	dataBlob := c.Body()

	decryptedData, err := utils.DecryptData(dataBlob, utils.PrivateKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Ошибка расшифровки данных")
	}

	var beaconData map[string]interface{}
	if err := json.Unmarshal(decryptedData, &beaconData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Ошибка в формате данных")
	}

	log.Printf("Полученные данные: %+v", beaconData)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Данные успешно приняты",
	})
}
