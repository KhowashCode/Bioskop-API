package database

import (
	"Bioskop-API/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	defer tableRegister()
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load dotenv!")
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=bioskop port=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))

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
