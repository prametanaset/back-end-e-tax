package main

import (
	"back-end-e-tax/config"
	"back-end-e-tax/internal/route"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Load config
	config.LoadEnv()

	// Setup routes
	route.SetupRoutes(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
