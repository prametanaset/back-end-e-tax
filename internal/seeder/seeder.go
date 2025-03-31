package seeder

import (
	"back-end-e-tax/config"
	"back-end-e-tax/internal/models"
	"log"
)

// SeedData ฟังก์ชันสำหรับเติมข้อมูลตัวอย่าง
func SeedData() {
	// ตรวจสอบว่า customer มีข้อมูลหรือไม่
	var count int64
	config.DB.Model(&models.Customer{}).Count(&count)
	if count == 0 {
		// สร้างข้อมูลตัวอย่างลูกค้า
		customers := []models.Customer{
			{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "johndoe@example.com",
				Phone:     "123-456-7890",
			},
			{
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "janedoe@example.com",
				Phone:     "098-765-4321",
			},
		}

		for _, customer := range customers {
			if err := config.DB.Create(&customer).Error; err != nil {
				log.Fatalf("Error inserting customer: %v", err)
			}
		}
		log.Println("Customer data seeded successfully!")
	}

	// ตรวจสอบว่า product มีข้อมูลหรือไม่
	config.DB.Model(&models.Product{}).Count(&count)
	if count == 0 {
		// สร้างข้อมูลตัวอย่างสินค้า
		products := []models.Product{
			{
				Name:        "Sample Product 1",
				Description: "This is a sample product 1",
				Price:       19.99,
				Stock:       100,
			},
			{
				Name:        "Sample Product 2",
				Description: "This is a sample product 2",
				Price:       29.99,
				Stock:       50,
			},
		}

		for _, product := range products {
			if err := config.DB.Create(&product).Error; err != nil {
				log.Fatalf("Error inserting product: %v", err)
			}
		}
		log.Println("Product data seeded successfully!")
	}
}
