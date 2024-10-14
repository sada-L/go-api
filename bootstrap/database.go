package bootstrap

import (
	"fmt"
	"go-server/internal/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database!", zap.Error(err))
	}

	return db
}
