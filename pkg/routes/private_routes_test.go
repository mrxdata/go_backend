package routes

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPrivateRoutes(t *testing.T) {
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}

	dataString := `{"id": "00000000-0000-0000-0000-000000000000"}`

	// Create Session

	tests := []struct {
		description   string
		route         string
		method        string
		tokenString   string
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "test1",
			route:         "/api/v1/chat",
			method:        "DELETE",
			body:          nil,
			expectedError: false,
			expectedCode:  404,
		},
		{
			description:   "test2",
			route:         "/api/v1/chat",
			method:        "DELETE",
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  404,
		},
		{
			description:   "test3",
			route:         "/api/v1/chat",
			method:        "DELETE",
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  404,
		},
	}

	app := fiber.New()

	PrivateRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Authorization", test.tokenString)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, fiber.TestConfig{
			Timeout:       100 * time.Millisecond,
			FailOnTimeout: false, // If false, do not discard response
		})

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
