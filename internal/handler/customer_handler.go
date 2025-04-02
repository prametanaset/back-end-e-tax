package handler

import (
	"back-end-e-tax/config"
	"back-end-e-tax/internal/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetCustomers ดึงข้อมูลลูกค้าทั้งหมด
func GetCustomers(c *fiber.Ctx) error {
	var customers []models.Customer
	// var customersRes []models.CustomerResponse

	// ดึงข้อมูลลูกค้าทั้งหมดจากฐานข้อมูล
	if err := config.DB.Find(&customers).Error; err != nil {
		log.Printf("Error fetching customers: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching customers",
		})
	}

	// ส่งกลับข้อมูลลูกค้า
	return c.Status(fiber.StatusOK).JSON(customers)
}

// UpdateCustomer updates an existing customer record
func UpdateCustomer(c *fiber.Ctx) error {
	var customer models.Customer

	// Bind the incoming JSON request body to the customer struct
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Retrieve the existing customer from the database
	var existingCustomer models.Customer
	if err := config.DB.First(&existingCustomer, customer.ID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	// Update the customer details
	customer.UpdatedAt = time.Now()

	// Save the updated customer record in the database
	if err := config.DB.Save(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update customer",
		})
	}

	// Return the updated customer object
	return c.Status(fiber.StatusOK).JSON(customer)
}

