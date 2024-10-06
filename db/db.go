package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	. "go-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=89.110.53.87 user=postgres password=postgres dbname=postgres port=5533 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Автоматическая миграция таблиц
	database.AutoMigrate(&User{})

	DB = database
}
