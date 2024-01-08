package config

import (
	"time"

	"github.com/spf13/viper"
)

// Contains all configuration variables
type Config struct {
	HTMXPort     string `mapstructure:"HTMX_PORT"`
	TemplatePath string `mapstructure:"TEMPLATE_PATH"`

	HTTPPort string `mapstructure:"HTTP_PORT"`
	WSPort   string `mapstructure:"WS_PORT"`
	GRPCPort string `mapstructure:"GRPC_PORT"`

	HTTPDocsPath string `mapstructure:"HTTP_DOCS_PATH"`

	DatabaseURL string `mapstructure:"DATABASE_URL"`

	RedisPort string `mapstructure:"REDIS_PORT"`
	RedisUser string `mapstructure:"REDIS_USER"`
	RedisPass string `mapstructure:"REDIS_PASS"`

	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`

	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func NewConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	// Override variables from file with the environmet variables
	viper.AutomaticEnv()
	c := Config{}
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
