package main

import (
	"back-end-e-tax/config"
	"back-end-e-tax/internal/routes"
	"back-end-e-tax/internal/seeder"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Load config
	config.LoadEnv()
	config.ConnectDB()

	// เติมข้อมูลตัวอย่างในฐานข้อมูล
	seeder.SeedData()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}


