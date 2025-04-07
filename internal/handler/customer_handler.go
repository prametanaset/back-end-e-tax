package handler

import (
	"back-end-e-tax/config"
	"back-end-e-tax/internal/models"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// --------------------------------------------------------------------------------------------------------------------
// GetCustomers ดึงข้อมูลลูกค้าทั้งหมด
func GetCustomers(c *fiber.Ctx) error {
	var customers []models.Customer
	// var customersRes []models.CustomerResponse

	// ใช้ SQL Query ดึงข้อมูลและแมปลง struct โดยตรง
	err := config.DB.Raw("SELECT * FROM customers ORDER BY updated_at DESC").Scan(&customers).Error
	if err != nil {
		log.Printf("Error fetching customers: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching customers",
		})
	}

	// ส่งกลับข้อมูลลูกค้า
	return c.Status(fiber.StatusOK).JSON(customers)
}

// --------------------------------------------------------------------------------------------------------------------
// UpdateCustomer updates an existing customer record
func UpdateCustomer(c *fiber.Ctx) error {
	// รับค่า ID จาก URL และแปลงเป็นตัวเลข
	log.Printf("fetching customers: %s", c.Body())
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    var customer models.Customer

    // ค้นหาสินค้าในฐานข้อมูล
    if err := config.DB.First(&customer, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    // อ่านข้อมูลที่ส่งมา
    var updateData models.Customer
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // อัปเดตข้อมูล
    customer.FirstName = updateData.FirstName
    customer.LastName = updateData.LastName
    customer.Email = updateData.Email
    customer.Address = updateData.Address
    customer.Phone = updateData.Phone
    customer.TaxIdNo = updateData.TaxIdNo
    customer.UpdatedAt = time.Now()

    // อัปเดตในฐานข้อมูล
    if err := config.DB.Model(&customer).Updates(updateData).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update customer"})
    }

    return c.JSON(fiber.Map{"message": "customer updated successfully", "customer": customer})
}
// --------------------------------------------------------------------------------------------------------------------

func CreateCustomer(c *fiber.Ctx) error {
	var  customer models.Customer

		// แปลง JSON ที่รับมาเป็น struct
	if err := c.BodyParser(&customer); err != nil {
		log.Printf("❌ Error parsing customer: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// ตรวจสอบข้อมูลเบื้องต้น
	if customer.FirstName == "" || customer.LastName == "" || customer.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required fields: FirstName, LastName, Email",
		})
	}

	// บันทึกลงฐานข้อมูล
	if err := config.DB.Create(&customer).Error; err != nil {
		log.Printf("❌ Error inserting customer: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create customer",
		})
	}

	// ส่งข้อมูลลูกค้าที่เพิ่มกลับไป
	return c.Status(fiber.StatusCreated).JSON(customer)

}

// --------------------------------------------------------------------------------------------------------------------
// DeleteCustomer ลบข้อมูลลูกค้าจากฐานข้อมูล
func DeleteCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id")) // รับค่า ID จาก URL
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var customer models.Customer

	// ค้นหาลูกค้าในฐานข้อมูล
	if err := config.DB.First(&customer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve customer"})
	}

	// ลบข้อมูลลูกค้า
	if err := config.DB.Delete(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete customer"})
	}

	// ส่งผลลัพธ์เมื่อลบสำเร็จ
	return c.JSON(fiber.Map{"message": "Customer deleted successfully"})
}


