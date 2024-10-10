package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-api/configs"
	"go-api/services"
	"go-api/utils"
	"net/http"
	"strings"
)

type Handlers struct {
	UserHandler
	ImageHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		UserHandler:  &userHandler{s.User},
		ImageHandler: &imageHandler{s.Image},
	}
}

func SetDefault(e *echo.Echo) {
	utils.SetHTMLTemplateRenderer(e)

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "data", configs.Auth0Config)
	})
	e.GET("/healthcheck", HealthCheckHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func SetApi(e *echo.Echo, h *Handlers, m echo.MiddlewareFunc) {
	g := e.Group("/api/v1")
	//g.Use(m)

	// User
	g.GET("/user", h.UserHandler.GetUsers)
	g.POST("/user", h.UserHandler.CreateUser)
	g.PUT("/user", h.UserHandler.UpdateUser)
	g.DELETE("/user/:id", h.UserHandler.DeleteUserById)

	//image
	g.GET("/image/:id", h.ImageHandler.GetImage)
	g.POST("/image", h.ImageHandler.UploadImage)
	g.DELETE("/image/:id", h.ImageHandler.DeleteImageById)
}

func Echo() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))

	return e
}
