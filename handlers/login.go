package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go-api/configs"
	"go-api/logger"
	"go-api/utils"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

// LoginHandler
// @Summary Авторизация.
// @Description Авторизация по логину и паролю.
// @Tags Auth
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {string} string "Your token"
// @Failure 401 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /login [post]
func LoginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "admin" || password != "123" {
		logger.Error("failed to authorized", zap.Error(echo.ErrUnauthorized))
		return c.JSON(http.StatusUnauthorized, utils.Error{Message: "failed to authorized"})
	}
	isAdmin := true

	claims := &configs.JwtClaims{
		Name:  username,
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		logger.Error("failed to encode JWT", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": t})
}
