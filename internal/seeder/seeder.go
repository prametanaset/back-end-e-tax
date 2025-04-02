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
    Price: 7858.82,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1001",
    Name: "เส้นใหญ่",
    Price: 2839.22,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1002",
    Name: "น้ำปลาร้า",
    Price: 3246.8,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1003",
    Name: "ข้าวเหนียว",
    Price: 5814.04,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1004",
    Name: "น้ำดื่ม",
    Price: 379.5,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1005",
    Name: "แป้งมัน",
    Price: 2312.87,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1006",
    Name: "นมข้นจืด",
    Price: 7299.78,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1007",
    Name: "ขนมปังกรอบ",
    Price: 2789.58,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1008",
    Name: "แยมผลไม้",
    Price: 59.9,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1009",
    Name: "ฟองเต้าหู้",
    Price: 9809.1,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1010",
    Name: "น้ำอัดลม",
    Price: 1146.02,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1011",
    Name: "ซอสมะเขือเทศ",
    Price: 9297.87,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1012",
    Name: "วุ้นเส้น",
    Price: 1981.58,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1013",
    Name: "หม้อหุงข้าว",
    Price: 9834.17,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1014",
    Name: "ช็อกโกแลต",
    Price: 1171.32,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1015",
    Name: "น้ำส้มสายชู",
    Price: 4410.5,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1016",
    Name: "แชมพู",
    Price: 3230.16,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1017",
    Name: "เส้นเล็ก",
    Price: 1159.56,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1018",
    Name: "ข้าวโพดกระป๋อง",
    Price: 5600.62,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1019",
    Name: "ขนมปัง",
    Price: 4311.0,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1020",
    Name: "เตาไฟฟ้า",
    Price: 3288.07,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1021",
    Name: "กาแฟสำเร็จรูป",
    Price: 7118.53,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1022",
    Name: "ไส้กรอก",
    Price: 3786.81,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1023",
    Name: "บะหมี่กึ่งสำเร็จรูป",
    Price: 6326.25,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1024",
    Name: "น้ำมันมะกอก",
    Price: 4093.25,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1025",
    Name: "โจ๊กซอง",
    Price: 3997.52,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1026",
    Name: "เห็ดหอมแห้ง",
    Price: 918.34,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1027",
    Name: "ซอสหอยนางรม",
    Price: 2841.9,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1028",
    Name: "ปลาทูน่ากระป๋อง",
    Price: 8758.26,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1029",
    Name: "มันฝรั่งทอด",
    Price: 625.51,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1030",
    Name: "ข้าวโอ๊ต",
    Price: 7072.28,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1031",
    Name: "แครกเกอร์",
    Price: 5418.0,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1032",
    Name: "น้ำเต้าหู้",
    Price: 3072.26,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1033",
    Name: "ถั่วลิสง",
    Price: 9007.96,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1034",
    Name: "น้ำผึ้ง",
    Price: 3421.48,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1035",
    Name: "ขวดน้ำ",
    Price: 7815.72,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1036",
    Name: "ซีอิ๊วขาว",
    Price: 4585.32,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1037",
    Name: "นมข้นหวาน",
    Price: 1876.1,
    Vat: false,
    VatRate: 0,
  },
  {
    ProductCode: "PROD-1038",
    Name: "ไมโครเวฟ",
    Price: 3063.09,
    Vat: true,
    VatRate: 7,
  },
  {
    ProductCode: "PROD-1039",
    Name: "ตู้เย็น",
    Price: 8025.05,
    Vat: false,
    VatRate: 0,
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
