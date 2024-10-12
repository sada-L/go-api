package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"go-api/configs"
	"go-api/logger"
	"go.uber.org/zap"
	"time"
)

func GenerateTokens(username string) (string, string, error) {
	accessClaims := &configs.JwtClaims{
		Username:  username,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(configs.JwtSecret)
	if err != nil {
		logger.Error("failed to encode JWT", zap.Error(err))
		return "", "", err
	}

	refreshClaims := configs.JwtClaims{
		Username:  username,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(configs.JwtSecret)
	if err != nil {
		logger.Error("failed to encode JWT", zap.Error(err))
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
