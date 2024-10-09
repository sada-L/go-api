package main

import (
	database "go-api/db"
	_ "go-api/docs"
	"go-api/handlers"
	"go-api/middlewares"
	"go-api/services"
	"go-api/stores"
	"log"
	"os"
)

// @title GO API
// @version 1.0
// @description Server for a user management API.

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	e := handlers.Echo()

	s := stores.New(db)
	ss := services.New(s)
	h := handlers.New(ss)

	jwtCheck, err := middlewares.JwtMiddleware()
	if err != nil {
		log.Fatal("failed to set JWT middleware")
	}

	handlers.SetDefault(e)
	handlers.SetApi(e, h, jwtCheck)

	if err := e.Start(os.Getenv("SERVER_ADDRESS")); err != nil {
		log.Fatal("Error in server start", err)
	}
}
