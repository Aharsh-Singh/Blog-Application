package main

import (
	"fmt"
	"myapp/Routes"
	"github.com/joho/godotenv"
	"log"
	"myapp/models"
	"myapp/config"
)
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := config.ConnectDatabase()
	db.AutoMigrate(&models.User{})
	r := Routes.SetupRouter()
    fmt.Println("Server is running on port 8080")
    r.Run(":8080")
}
