package config

import (
	"back-end-e-tax/internal/models"
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

    // ทำการ migrate schema
    err = DB.AutoMigrate(&models.Invoice{},&models.Customer{}, &models.Product{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
    log.Println("Database migrated successfully!")
}
