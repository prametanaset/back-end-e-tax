package handler

import (
	"back-end-e-tax/config"
	"back-end-e-tax/internal/models"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// --------------------------------------------------------------------------------------------------------------------
// GetProducts ดึงข้อมูลสินค้าทั้งหมดจากฐานข้อมูล
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	// ใช้ SQL Query ดึงข้อมูลและแมปลง struct โดยตรง
	err := config.DB.Raw("SELECT * FROM products ORDER BY updated_at DESC").Scan(&products).Error
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching products",
		})
	}

	// ส่งข้อมูลกลับเป็น JSON
	return c.Status(fiber.StatusOK).JSON(products)
}

// --------------------------------------------------------------------------------------------------------------------
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


// --------------------------------------------------------------------------------------------------------------------
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
// --------------------------------------------------------------------------------------------------------------------
// UpdateProduct อัปเดตข้อมูลสินค้า
func UpdateProduct(c *fiber.Ctx) error {
    // รับค่า ID จาก URL และแปลงเป็นตัวเลข
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    var product models.Product

    // ค้นหาสินค้าในฐานข้อมูล
    if err := config.DB.First(&product, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    // อ่านข้อมูลที่ส่งมา
    var updateData models.Product
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // อัปเดตข้อมูล
    product.Name = updateData.Name
    product.Price = updateData.Price
    product.Vat = updateData.Vat
    product.VatRate = updateData.VatRate
    product.UpdatedAt = time.Now()

    // อัปเดตในฐานข้อมูล
    if err := config.DB.Model(&product).Updates(updateData).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
    }

    return c.JSON(fiber.Map{"message": "Product updated successfully", "product": product})
}

// --------------------------------------------------------------------------------------------------------------------
// DeleteProduct ลบข้อมูลสินค้าจากฐานข้อมูล
func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id")) // รับค่า ID จาก URL
	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

	var product models.Product

	// ค้นหาสินค้าในฐานข้อมูล
	if err := config.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// ลบข้อมูลสินค้า
	if err := config.DB.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
	}

	// ส่งผลลัพธ์เมื่อลบสำเร็จ
	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}

