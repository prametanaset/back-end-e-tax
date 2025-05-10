package main

import (
	"back-end-e-tax/config"
	"back-end-e-tax/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
		// AllowCredentials: true, // ✅ ใช้ได้เมื่อ AllowOrigins ไม่ใช่ '*'
	}))

	// Load config
	config.LoadEnv()
	config.ConnectDB()

	// เติมข้อมูลตัวอย่างในฐานข้อมูล
	// seeder.SeedData()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
