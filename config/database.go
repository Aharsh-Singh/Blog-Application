package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"log"
	"myapp/models"
)

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to Aiven PostgreSQL successfully 🚀")
	models.DB = db
	return db
}
