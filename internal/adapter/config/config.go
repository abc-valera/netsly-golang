package config

import (
	"errors"
	"os"
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
)

var Mode string

const (
	DevelopmentMode string = "dev"
	ProductionMode  string = "prod"
)

const (
	ModeKey = "MODE"

	HTMXPortKey     = "HTMX_PORT"
	TemplatePathKey = "TEMPLATE_PATH"

	HTTPPortKey     = "HTTP_PORT"
	HTTPDocsPathKey = "HTTP_DOCS_PATH"
	WSPortKey       = "WS_PORT"

	GRPCPortKey = "GRPC_PORT"

	PostgresURLKey = "POSTGRES_URL"

	RedisPortKey = "REDIS_PORT"
	RedisUserKey = "REDIS_USER"
	RedisPassKey = "REDIS_PASS"

	AccessTokenDurationKey  = "ACCESS_TOKEN_DURATION"
	RefreshTokenDurationKey = "REFRESH_TOKEN_DURATION"

	EmailSenderAddressKey  = "EMAIL_SENDER_ADDRESS"
	EmailSenderPasswordKey = "EMAIL_SENDER_PASSWORD"
)

// Contains all configuration variables
type Config struct {
	HTMXPort     string
	TemplatePath string

	HTTPPort     string
	HTTPDocsPath string
	WSPort       string

	GRPCPort string

	DatabaseURL string

	RedisPort string
	RedisUser string
	RedisPass string

	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration

	EmailSenderAddress  string
	EmailSenderPassword string
}

func NewConfig(configPath string) (Config, error) {
	mode := os.Getenv("MODE")
	if mode == "dev" || mode == "prod" {
		Mode = mode
	} else {
		return Config{}, coderr.NewInternal(errors.New("'MODE' environmental variable is not set"))
	}

	config := Config{}

	config.HTMXPort = os.Getenv(HTMXPortKey)
	config.TemplatePath = os.Getenv(TemplatePathKey)

	config.HTTPPort = os.Getenv(HTTPPortKey)
	config.HTTPDocsPath = os.Getenv(HTTPDocsPathKey)
	config.WSPort = os.Getenv(WSPortKey)

	config.GRPCPort = os.Getenv(GRPCPortKey)

	config.DatabaseURL = os.Getenv(PostgresURLKey)

	config.RedisPort = os.Getenv(RedisPortKey)
	config.RedisUser = os.Getenv(RedisUserKey)
	config.RedisPass = os.Getenv(RedisPassKey)

	if accessTokenDurationEnv := os.Getenv(AccessTokenDurationKey); accessTokenDurationEnv != "" {
		accessTokenDuration, err := time.ParseDuration(accessTokenDurationEnv)
		if err != nil {
			return Config{}, coderr.NewInternal(errors.New("invalid value for 'ACCESS_TOKEN_DURATION' environmental variable"))
		}
		config.AccessTokenDuration = accessTokenDuration
	}

	if refreshTokenDurationEnv := os.Getenv(RefreshTokenDurationKey); refreshTokenDurationEnv != "" {
		refreshTokenDuration, err := time.ParseDuration(refreshTokenDurationEnv)
		if err != nil {
			return Config{}, coderr.NewInternal(errors.New("invalid value for 'REFRESH_TOKEN_DURATION' environmental variable"))
		}
		config.RefreshTokenDuration = refreshTokenDuration
	}

	config.EmailSenderAddress = os.Getenv(EmailSenderAddressKey)
	config.EmailSenderPassword = os.Getenv(EmailSenderPasswordKey)

	return config, nil
}
