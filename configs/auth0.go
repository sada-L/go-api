package configs

import (
	"github.com/auth0/go-jwt-middleware/v2/validator"
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

var Auth0Config = Auth0ConfigType{
	Domain:             "your-app-name.auth0.com",
	ClientID:           "123abcDEF456ghiJKL789mnoPQR",
	Audience:           []string{"https://your-api.example.com"},
	Issuer:             "https://your-app-name.auth0.com/",
	SignatureAlgorithm: validator.RS256,
	CacheDuration:      15 * time.Minute,
}
