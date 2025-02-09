package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPublicRoutes(t *testing.T) {
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}

	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "tg auth",
			route:         "/api/auth/telegram",
			expectedError: false,
			expectedCode:  405,
		},
		{
			description:   "1",
			route:         "/api/test",
			expectedError: false,
			expectedCode:  404,
		},
	}

	app := fiber.New()

	PublicRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, http.NoBody)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, fiber.TestConfig{
			Timeout:       100 * time.Millisecond,
			FailOnTimeout: false,
		})

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
