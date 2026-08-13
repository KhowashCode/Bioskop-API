package database

import (
	"Bioskop-API/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	defer tableRegister()
	dsn := "host=localhost user=postgres password=12345 dbname=bioskop port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to open database connection!")
	}

	DB = database

	fmt.Println("Success to open database connection...")
}

func tableRegister() {
	DB.AutoMigrate(&models.Bioskop{})
}
