package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB เชื่อมต่อฐานข้อมูล PostgreSQL
func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// เรียกใช้งาน Migration แยกไฟล์
	// if err := migration.Migrate(DB); err != nil {
	// 	log.Fatalf("❌ Failed to migrate: %v", err)
	// }

	// if err := seeder.Seed(DB); err != nil {
	// 	log.Fatal("seeding failed:", err)
	// }

	// log.Println("Database migrated successfully!")
}
