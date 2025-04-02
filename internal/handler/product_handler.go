package handler

import (
	"back-end-e-tax/config"
	"back-end-e-tax/internal/models"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetProducts ดึงข้อมูลสินค้าทั้งหมดจากฐานข้อมูล
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	// ใช้ SQL Query ดึงข้อมูลและแมปลง struct โดยตรง
	err := config.DB.Raw("SELECT * FROM products ORDER BY id DESC").Scan(&products).Error
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching products",
		})
	}

	// ส่งข้อมูลกลับเป็น JSON
	return c.Status(fiber.StatusOK).JSON(products)
}

// CreateProduct เพิ่มสินค้าใหม่ลงในฐานข้อมูล
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	// แปลง JSON ที่รับมาเป็น struct
	if err := c.BodyParser(&product); err != nil {
		log.Printf("Error parsing product: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// ถ้าไม่ได้ส่ง ProductCode มาให้ Generate ใหม่
    if product.ProductCode == "" {
        newCode, err := GenerateProductCode()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "Error generating product code",
            })
        }
        product.ProductCode = newCode
    }

	// บันทึกลงฐานข้อมูล
	if err := config.DB.Create(&product).Error; err != nil {
		log.Printf("Error inserting product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error inserting product",
		})
	}

	// ส่งข้อมูลสินค้าที่เพิ่มกลับไป
	return c.Status(fiber.StatusCreated).JSON(product)
}

func GenerateProductCode() (string, error) {
    var lastProduct models.Product
    var lastNumber int

    // ค้นหา Product ล่าสุดที่มี ProductCode สูงสุด
    result := config.DB.Order("id DESC").First(&lastProduct)
    if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return "", result.Error
    }

    // ถ้ามีสินค้าแล้ว ดึงเลขที่มากที่สุด
    if lastProduct.ProductCode != "" {
        _, err := fmt.Sscanf(lastProduct.ProductCode, "PROD-%04d", &lastNumber)
        if err != nil {
            return "", err
        }
    }

    // เพิ่มเลข +1 และสร้าง ProductCode ใหม่
    newCode := fmt.Sprintf("PROD-%04d", lastNumber+1)
    return newCode, nil
}


