package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"flutty_messenger/app/models"
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
	"sort"
	"strconv"
	"strings"
)

func NewTgValidator() *validator.Validate {
	validate := validator.New()

	validate.RegisterStructValidation(TgAuthStructValidation, models.TgAuthData{})

	return validate
}

func TgAuthStructValidation(sl validator.StructLevel) {
	authData := sl.Current().Interface().(models.TgAuthData)

	dataCheckString := createDataCheckString(&authData)

	secret := os.Getenv("TELEGRAM_BOT_SECRET")
	if secret == "" {
		sl.ReportError(secret, "TELEGRAM_BOT_SECRET", "TELEGRAM_BOT_SECRET", "missing_secret", "")
		return
	}

	secretKey := sha256.Sum256([]byte(secret))

	expectedHash := generateHash(dataCheckString, secretKey)

	if expectedHash != authData.Hash {
		sl.ReportError(authData.Hash, "hash", "hash", "invalid_hash", "")
	}
}

func createDataCheckString(data *models.TgAuthData) string {
	keys := []string{"auth_date", "first_name", "id", "last_name", "photo_url", "username"}
	sort.Strings(keys)

	var parts []string
	for _, key := range keys {
		var value string
		switch key {
		case "auth_date":
			value = strconv.Itoa(data.AuthDate)
		case "first_name":
			value = data.FirstName
		case "id":
			value = strconv.Itoa(data.Id)
		case "last_name":
			value = data.LastName
		case "photo_url":
			value = data.PhotoURL
		case "username":
			value = data.Username
		}
		if value != "" {
			parts = append(parts, fmt.Sprintf("%s=%s", key, value))
		}
	}
	return strings.Join(parts, "\n")
}

func generateHash(dataCheckString string, secretKey [32]byte) string {
	h := hmac.New(sha256.New, secretKey[:])
	h.Write([]byte(dataCheckString))
	calculatedHash := hex.EncodeToString(h.Sum(nil))
	return calculatedHash
}

func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}
