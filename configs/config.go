package configs

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
)

type JwtClaims struct {
	Username  string `json:"username"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func NewAccessConf() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtClaims)
		},
		SigningKey: JwtSecret,
	}
}
