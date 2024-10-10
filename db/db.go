package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-api/docs"
	"go-api/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func init() {
	err := godotenv.Load("./env.env")
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
}

func New() (*gorm.DB, error) {
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PUBLIC_PORT")
	swaggerHost := fmt.Sprintf("%s:%s", serverHost, serverPort)
	if swaggerHost != "" {
		docs.SwaggerInfo.Host = swaggerHost
	}

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database!", zap.Error(err))
	}

	return db, err
}
