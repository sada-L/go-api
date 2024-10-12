package handlers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go-api/configs"
	"go-api/logger"
	"go-api/utils"
	"go.uber.org/zap"
	"net/http"
)

// RefreshHandler
// @Summary Обновление доступа.
// @Description Обновляет токен досутпа с помощью токена обновления.
// @Tags Auth
// @Produce json
// @Security Bearer Authentication
// @Param Authorization header string true "'Bearer refresh_token'"
// @Success 200 {object} TokenResponse "JWT tokens"
// @Failure 401 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /refresh [post]
func RefreshHandler(c echo.Context) error {
	refreshToken, e := ValidateTokenType(c)
	if e == nil {
		return c.JSON(http.StatusUnauthorized, utils.Error{Message: "invalid token"})
	}

	token, err := jwt.ParseWithClaims(refreshToken, &configs.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return configs.JwtSecret, nil
	})

	if err != nil {
		logger.Error("failed to validate refresh token", zap.Error(err))
		return c.JSON(http.StatusUnauthorized, utils.Error{Message: "failed to validate refresh token"})
	}

	if claims, ok := token.Claims.(*configs.JwtClaims); ok && token.Valid {
		newAccessToken, _, err := utils.GenerateTokens(claims.Username)
		if err != nil {
			logger.Error("failed to generate new access token", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, utils.Error{Message: "failed to generate new access token"})
		}

		return c.JSON(http.StatusOK, TokenResponse{
			AccessToken:  newAccessToken,
			RefreshToken: refreshToken,
		})
	} else {
		logger.Error("Invalid refresh token", zap.Error(err))
		return c.JSON(http.StatusUnauthorized, utils.Error{Message: "Invalid refresh token"})
	}
}

func ValidateTokenType(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return "", errors.New("invalid authorization header")
	}

	authToken := authHeader[7:]
	token, err := jwt.ParseWithClaims(authToken, &configs.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return configs.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims := token.Claims.(*configs.JwtClaims)
	if claims.TokenType != "access" {
		return authToken, errors.New("invalid token type")
	}

	return authToken, nil
}
