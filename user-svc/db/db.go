package db

import (
	"log"
	"main/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL") 
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set.")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//fmt.Println("✅ Connected to PostgreSQL successfully!")
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
}