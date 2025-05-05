package seeder

import (
	"back-end-e-tax/pkg/model"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// 1. Address
	address := model.Address{}
	db.FirstOrCreate(&address, model.Address{
		Street: "99 Seed Rd", District: "Bang Rak", Province: "Bangkok",
	})
	fmt.Println("Address ID:", address.ID)

	// 2. Tenant
	tenant := model.Tenant{}
	db.First(&tenant, "tax_id = ?", "1234567890123")
	if tenant.ID == 0 {
		tenant = model.Tenant{
			TaxID:     "1234567890123",
			LegalName: "Example Corp", BranchCode: "0001",
			AddressID: address.ID,
			LogoURL:   "https://logo.png",
		}
		db.Create(&tenant)
	}
	fmt.Println("Tenant ID:", tenant.ID)

	// 3. User
	user := model.User{}
	db.First(&user, "username = ?", "admin")
	if user.ID == 0 {
		user = model.User{
			TenantID: tenant.ID,
			Username: "admin",
			Password: "$2a$12$examplepass",
			FullName: "Admin User",
			Email:    "admin@corp.com",
			IsActive: true,
		}
		db.Create(&user)
	}
	fmt.Println("User ID:", user.ID)

	// 4. Role
	role := model.Role{}
	db.First(&role, "name = ?", "admin")
	if role.ID == 0 {
		role = model.Role{Name: "admin"}
		db.Create(&role)
	}
	fmt.Println("Role ID:", role.ID)

	// 5. Attach Role
	var roles []model.Role
	db.Model(&user).Association("Roles").Find(&roles)
	found := false
	for _, r := range roles {
		if r.ID == role.ID {
			found = true
			break
		}
	}
	if !found {
		db.Model(&user).Association("Roles").Append(&role)
		fmt.Println("Role assigned to user")
	}

	// 6. Buyer
	buyer := model.Buyer{}
	db.First(&buyer, "tax_id = ?", "9999999999999")
	if buyer.ID == 0 {
		buyer = model.Buyer{
			Name: "บจก. ลูกค้า", TaxID: "9999999999999",
			AddressID: address.ID, ContactPerson: "คุณลูกค้า",
			Email: "client@example.com",
		}
		db.Create(&buyer)
	}
	fmt.Println("Buyer ID:", buyer.ID)

	// 7. VAT Category
	vat := model.VatCategory{}
	db.First(&vat, "description = ?", "VAT 7%")
	if vat.ID == 0 {
		vat = model.VatCategory{Description: "VAT 7%", VatRate: 7}
		db.Create(&vat)
	}
	fmt.Println("VAT ID:", vat.ID)

	// 8. Product 50 รายการ
	thaiProducts := []string{
		"ปากกาลูกลื่น", "ดินสอกด", "ยางลบ", "สมุดบันทึก", "แฟ้มเอกสาร",
		"กระดาษถ่ายเอกสาร", "คลิปหนีบกระดาษ", "เครื่องเย็บกระดาษ", "กรรไกร", "เทปใส",
		"เทปกาวสองหน้า", "ไม้บรรทัด", "คัตเตอร์", "กล่องเอกสาร", "โฟลเดอร์พลาสติก",
		"แผ่นซีดี", "ปากกาเน้นข้อความ", "กระดานไวท์บอร์ด", "หมึกเติมปากกา", "กาวน้ำ",
		"แม็กเย็บกระดาษ", "แม็กเย็บเอกสาร", "กระดาษโน้ต", "แผ่นรองเขียน", "ปากกาเจล",
		"แฟ้มหนีบ", "สติ๊กเกอร์กระดาษ", "ซองจดหมาย", "กระดาษรายงาน", "กระดาษการ์ดสี",
		"แฟ้มสันรูด", "แฟ้มกระดุม", "แฟ้มเจาะ", "ปากกาไวท์บอร์ด", "น้ำยาเช็ดกระดาน",
		"หมึกพิมพ์", "ตลับหมึก", "เทปผ้า", "กระดาษกาว", "กระดาษทราย",
		"แผ่นรองเมาส์", "ปากกาเมจิก", "แฟ้มใส่เอกสาร", "ที่เย็บกระดาษไฟฟ้า", "กล่องจดหมาย",
		"เครื่องเจาะกระดาษ", "กล่องเก็บของ", "ถุงซิปล็อค", "ไม้หนีบเอกสาร", "เครื่องคิดเลข",
	}

	for i, name := range thaiProducts {
		sku := fmt.Sprintf("T%03d", i+1)
		var prod model.Product
		db.First(&prod, "sku = ? AND tenant_id = ?", sku, tenant.ID)
		if prod.ID == 0 {
			price := float64(15 + i*5)
			prod = model.Product{
				TenantID:      tenant.ID,
				SKU:           sku,
				Name:          name,
				UOM:           "ชิ้น",
				PriceStandard: price,
				VatCategoryID: vat.ID,
			}
			db.Create(&prod)
			fmt.Printf("Created Product %s: %s (%.2f฿)\n", sku, name, price)
		}
	}

	// 9. เลือกสินค้าตัวแรกมาใช้ใน Invoice Item
	var firstProduct model.Product
	db.Where("tenant_id = ?", tenant.ID).Order("id asc").First(&firstProduct)

	// 10. Invoice
	invoice := model.Invoice{}
	db.First(&invoice, "invoice_number = ?", "INV-0001")
	if invoice.ID == 0 {
		invoice = model.Invoice{
			TenantID: tenant.ID, BuyerID: buyer.ID, InvoiceNumber: "INV-0001",
			IssueDate: time.Now(), CurrencyCode: "THB",
			SubTotal:     firstProduct.PriceStandard,
			VatBaseTotal: firstProduct.PriceStandard,
			VatTotal:     firstProduct.PriceStandard * 0.07,
			GrandTotal:   firstProduct.PriceStandard * 1.07,
			Status:       "created",
		}
		db.Create(&invoice)
	}
	fmt.Println("Invoice ID:", invoice.ID)

	// 11. Invoice Item
	item := model.InvoiceItem{}
	db.First(&item, "invoice_id = ? AND product_id = ?", invoice.ID, firstProduct.ID)
	if item.ID == 0 {
		item = model.InvoiceItem{
			InvoiceID:   invoice.ID,
			ProductID:   firstProduct.ID,
			Description: firstProduct.Name,
			Qty:         1,
			UOM:         firstProduct.UOM,
			UnitPrice:   firstProduct.PriceStandard,
			VatRate:     7,
			VatAmt:      firstProduct.PriceStandard * 0.07,
			LineTotal:   firstProduct.PriceStandard * 1.07,
		}
		db.Create(&item)
	}
	fmt.Println("Invoice Item ID:", item.ID)

	// 12. Audit Log
	log := model.AuditLog{}
	db.First(&log, "reference_id = ?", invoice.InvoiceNumber)
	if log.ID == 0 {
		log = model.AuditLog{
			UserID:      &user.ID,
			TenantID:    tenant.ID,
			Module:      "invoice",
			Action:      "create",
			Level:       "INFO",
			ReferenceID: invoice.InvoiceNumber,
			Message:     "สร้างใบกำกับภาษี",
			DataAfter:   datatypes.JSON([]byte(fmt.Sprintf(`{"invoice_number":"%s"}`, invoice.InvoiceNumber))),
			IPAddress:   "127.0.0.1",
			UserAgent:   "Seeder",
			CreatedAt:   time.Now(),
		}
		db.Create(&log)
	}
	fmt.Println("Audit Log ID:", log.ID)

	return nil
}
