package config

import (
	"errors"
	"os"
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/global"
)

// Contains all configuration variables.
//
// Values are loaded from environmental variables.
type Config struct {
	WebAppPort         string
	WebAppTemplatePath string

	JsonRestApiPort       string
	JsonRestApiStaticPath string

	GRPCApiPort string

	PosrgresURL string

	RedisPort string
	RedisUser string
	RedisPass string

	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration

	EmailSenderAddress  string
	EmailSenderPassword string
}

func NewConfig() (Config, error) {
	mode := os.Getenv("MODE")
	switch mode {
	case "dev":
		global.Mode = global.ModeDevelopment
	case "prod":
		global.Mode = global.ModeProduction
	default:
		return Config{}, coderr.NewInternal(errors.New("'MODE' environmental variable is not set"))
	}

	config := Config{}

	config.WebAppPort = os.Getenv("WEB_APP_PORT")
	config.WebAppTemplatePath = os.Getenv("WEB_APP_TEMPLATE_PATH")

	config.JsonRestApiPort = os.Getenv("JSON_REST_API_PORT")
	config.JsonRestApiStaticPath = os.Getenv("JSON_REST_API_STATIC_PATH")

	config.GRPCApiPort = os.Getenv("GRPC_API_PORT")

	config.PosrgresURL = os.Getenv("POSTGRES_URL")

	config.RedisPort = os.Getenv("REDIS_PORT")
	config.RedisUser = os.Getenv("REDIS_USER")
	config.RedisPass = os.Getenv("REDIS_PASS")

	if accessTokenDurationEnv := os.Getenv("ACCESS_TOKEN_DURATION"); accessTokenDurationEnv != "" {
		accessTokenDuration, err := time.ParseDuration(accessTokenDurationEnv)
		if err != nil {
			return Config{}, coderr.NewInternal(errors.New("invalid value for 'ACCESS_TOKEN_DURATION' environmental variable"))
		}
		config.AccessTokenDuration = accessTokenDuration
	}

	if refreshTokenDurationEnv := os.Getenv("REFRESH_TOKEN_DURATION"); refreshTokenDurationEnv != "" {
		refreshTokenDuration, err := time.ParseDuration(refreshTokenDurationEnv)
		if err != nil {
			return Config{}, coderr.NewInternal(errors.New("invalid value for 'REFRESH_TOKEN_DURATION' environmental variable"))
		}
		config.RefreshTokenDuration = refreshTokenDuration
	}

	config.EmailSenderAddress = os.Getenv("EMAIL_SENDER_ADDRESS")
	config.EmailSenderPassword = os.Getenv("EMAIL_SENDER_PASSWORD")

	return config, nil
}
