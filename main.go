package main

import (
	"github.com/gin-gonic/gin"
	"go-server/api/route"
	"go-server/bootstrap"
	"go-server/internal/logger"
	"go.uber.org/zap"
	"log"
	"time"
)

// @title GO API
// @version 1.0
// @description Server for an image management API.

// @BasePath /
// @schemes http
func main() {
	err := logger.New()
	if err != nil {
		log.Fatal(err)
	}

	app := bootstrap.App()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

	e := gin.Default()

	route.Setup(app.Env, timeout, app.Database, e)

	logger.Fatal("failed to start server", zap.Error(e.Run(app.Env.ServerPort)))
}
