package handler

import (
	"back-end-e-tax/config"
	"back-end-e-tax/internal/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// GetCustomers ดึงข้อมูลลูกค้าทั้งหมด
func GetCustomers(c *fiber.Ctx) error {
	var customers []models.Customer
	var customersRes []models.CustomerResponse

	// ดึงข้อมูลลูกค้าทั้งหมดจากฐานข้อมูล
	if err := config.DB.Find(&customers).Error; err != nil {
		log.Printf("Error fetching customers: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching customers",
		})
	}

	// ส่งกลับข้อมูลลูกค้า
	return c.Status(fiber.StatusOK).JSON(customersRes)
}
