package configs

import (
	"fmt"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"os"
	"time"
)

type Auth0ConfigType struct {
	Domain             string
	ClientID           string
	Audience           []string
	Issuer             string
	SignatureAlgorithm validator.SignatureAlgorithm
	CacheDuration      time.Duration
}

func GetConf() *Auth0ConfigType {
	return &Auth0ConfigType{
		Domain:             fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PUBLIC_PORT")),
		ClientID:           "123abcDEF456ghiJKL789mnoPQR",
		Audience:           []string{"https://your-api.example.com"},
		Issuer:             "https://your-app-name.auth0.com/",
		SignatureAlgorithm: validator.RS256,
		CacheDuration:      15 * time.Minute,
	}
}
