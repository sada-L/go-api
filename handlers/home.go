package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Data struct {
	Host string
}

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "data", Data{Host: os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PUBLIC_PORT")})
}
