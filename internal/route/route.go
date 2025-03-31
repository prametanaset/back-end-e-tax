package route

import (
	"back-end-e-tax/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	// Example route
	v1.Get("/users", handler.GetUsers)
}
