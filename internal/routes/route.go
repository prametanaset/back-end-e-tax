package routes

import (
	"back-end-e-tax/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	// Example route
	v1.Get("/users", handler.GetUsers) 

	// customer
	v1.Get("/customers", handler.GetCustomers) 
	v1.Put("/customers", handler.UpdateCustomer) 

	// product
	v1.Get("/products", handler.GetProducts)
	v1.Post("/products", handler.CreateProduct)
	v1.Put("/products/:id", handler.UpdateProduct)
	v1.Delete("/products/:id", handler.DeleteProduct)
}
