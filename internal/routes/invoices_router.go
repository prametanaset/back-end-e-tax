package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"back-end-e-tax/internal/models"
)

func SetupInvoiceRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/invoices")

	api.Get("/", func(c *fiber.Ctx) error {
		var invoices []models.Invoice
		return c.JSON(invoices)
	})

	api.Post("/", func(c *fiber.Ctx) error {
		var invoice models.Invoice
		if err := c.BodyParser(&invoice); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		db.Create(&invoice)
		return c.JSON(invoice)
	})

	api.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var invoice models.Invoice
		if err := db.First(&invoice, "id = ?", id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Invoice not found"})
		}
		return c.JSON(invoice)
	})

	api.Put("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var invoice models.Invoice
		if err := db.First(&invoice, "id = ?", id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Invoice not found"})
		}
		if err := c.BodyParser(&invoice); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		db.Save(&invoice)
		return c.JSON(invoice)
	})

	api.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		db.Delete(&models.Invoice{}, "id = ?", id)
		return c.JSON(fiber.Map{"message": "Invoice deleted"})
	})
}
