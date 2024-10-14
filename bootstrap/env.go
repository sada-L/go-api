package bootstrap

import (
	"github.com/spf13/viper"
	"go-server/docs"
	"go-server/internal/logger"
	"go.uber.org/zap"
)

type Env struct {
	AppEnv                 string `json:"app_env"                   mapstructure:"APP_ENV"`
	ServerAddress          string `json:"server_address"            mapstructure:"SERVER_ADDRESS"`
	ServerPort             string `json:"server_port"               mapstructure:"SERVER_PORT"`
	PublicPort             string `json:"public_port"               mapstructure:"PUBLIC_PORT"`
	ContextTimeout         int    `json:"context_timeout"           mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `json:"db_host"                   mapstructure:"DB_HOST"`
	DBPort                 string `json:"db_port"                   mapstructure:"DB_PORT"`
	DBUser                 string `json:"db_user"                   mapstructure:"DB_USER"`
	DBPass                 string `json:"db_pass"                   mapstructure:"DB_PASS"`
	DBName                 string `json:"db_name"                   mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `json:"access_token_expiry_hour"  mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `json:"refresh_token_expiry_hour" mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `json:"access_token_secret"       mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `json:"refresh_token_secret"      mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("Can't find the file .env : ", zap.Error(err))
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		logger.Fatal("Environment can't be loaded: ", zap.Error(err))
	}

	if env.AppEnv == "development" {
		logger.Info("The App is running in development env")
	}
	docs.SwaggerInfo.Host = env.ServerAddress + env.PublicPort
	return &env
}
