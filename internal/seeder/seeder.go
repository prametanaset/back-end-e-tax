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
				Address:     "123 หมู่ที่ 16 ถ. มิตรภาพ ตำบลในเมือง อำเภอเมืองขอนแก่น ขอนแก่น 40002",
			},
			{
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "janedoe@example.com",
				Phone:     "098-765-4321",
				Address:     "ตำบล ศิลา อำเภอเมืองขอนแก่น ขอนแก่น 40000",
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
    ProductCode: "PROD-1000",
    Name: "ขนมขบเคี้ยว",
    StoreId: 1,
    Price: 7858.82,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1001",
    Name: "เส้นใหญ่",
    StoreId: 1,
    Price: 2839.22,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1002",
    Name: "น้ำปลาร้า",
    StoreId: 1,
    Price: 3246.8,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1003",
    Name: "ข้าวเหนียว",
    StoreId: 1,
    Price: 5814.04,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1004",
    Name: "น้ำดื่ม",
    StoreId: 1,
    Price: 379.5,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1005",
    Name: "แป้งมัน",
    StoreId: 1,
    Price: 2312.87,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1006",
    Name: "นมข้นจืด",
    StoreId: 1,
    Price: 7299.78,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1007",
    Name: "ขนมปังกรอบ",
    StoreId: 1,
    Price: 2789.58,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1008",
    Name: "แยมผลไม้",
    StoreId: 1,
    Price: 59.9,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1009",
    Name: "ฟองเต้าหู้",
    StoreId: 1,
    Price: 9809.1,
    Vat: true,
    VatRate: 7,
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
