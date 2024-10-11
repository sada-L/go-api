package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"go-api/configs"
	database "go-api/db"
	_ "go-api/docs"
	"go-api/handlers"
	"go-api/logger"
	"go-api/services"
	"go-api/stores"
	"go.uber.org/zap"
	"log"
	"os"
)

// @title GO API
// @version 1.0
// @description Server for a user management API.

// @BasePath /
// @schemes http
func main() {
	err := logger.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New()
	if err != nil {
		logger.Fatal("failed to connect to the database", zap.Error(err))
	}

	e := handlers.Echo()

	s := stores.New(db)
	ss := services.New(s)
	h := handlers.New(ss)

	jwtCheck := echojwt.WithConfig(configs.NewConf())

	handlers.SetDefault(e)
	handlers.SetApi(e, h, jwtCheck)

	port := os.Getenv("SERVER_LOCAL_PORT")
	logger.Fatal("failed to start server", zap.Error(e.Start(port)))
}
