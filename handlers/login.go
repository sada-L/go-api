package handlers

import (
	"github.com/labstack/echo/v4"
	"go-api/logger"
	"go-api/utils"
	"go.uber.org/zap"
	"net/http"
)

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// LoginHandler
// @Summary Авторизация.
// @Description Авторизация по логину и паролю.
// @Tags Auth
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} TokenResponse "JWT tokens"
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

	accessToken, refreshToken, err := utils.GenerateTokens(username)
	if err != nil {
		logger.Error("Failed to generate tokens", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: "Failed to generate tokens"})
	}

	return c.JSON(http.StatusOK, TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
