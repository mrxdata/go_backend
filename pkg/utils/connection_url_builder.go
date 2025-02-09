package utils

import (
	"fmt"
	"os"
)

func ConnectionURLBuilder(name string) (string, error) {
	var url string

	switch name {
	case "postgres":
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("PG_DB_HOST"),
			os.Getenv("PG_DB_PORT"),
			os.Getenv("PG_DB_USER"),
			os.Getenv("PG_DB_PASSWORD"),
			os.Getenv("PG_DB_NAME"),
			os.Getenv("PG_DB_SSL_MODE"),
		)
	case "redis":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	case "fiber":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		return "", fmt.Errorf("connection name '%v' is not supported", name)
	}

	return url, nil
}
