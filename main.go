package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	. "go-api/controllers"
	. "go-api/db"
	_ "go-api/docs"
	"log"
)

// @title GO API
// @version 1.0
// @description Server for a user management API.

// @host 89.110.53.87:5511
// @BasePath /

func registerRoutes(r *gin.Engine) {
	r.POST("/users", CreateUser)
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUserByID)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)
}

func main() {
	ConnectDatabase()
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Маршруты
	registerRoutes(r)

	// Запуск сервера
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
