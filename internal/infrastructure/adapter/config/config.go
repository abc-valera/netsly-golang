package config

import (
	"time"

	"github.com/spf13/viper"
)

// Contains all configuration variables
type Config struct {
	HTTPPort string `mapstructure:"HTTP_PORT"`
	WSPort   string `mapstructure:"WS_PORT"`
	GRPCPort string `mapstructure:"GRPC_PORT"`

	PostgreHost     string `mapstructure:"POSTGRES_HOST"`
	PostgrePort     string `mapstructure:"POSTGRES_PORT"`
	PostgreUser     string `mapstructure:"POSTGRES_USER"`
	PostgrePassword string `mapstructure:"POSTGRES_PASS"`
	PostgreName     string `mapstructure:"POSTGRES_NAME"`

	RedisPort string `mapstructure:"REDIS_PORT"`
	RedisUser string `mapstructure:"REDIS_USER"`
	RedisPass string `mapstructure:"REDIS_PASS"`

	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`

	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func InitConfig(configPath string) (*Config, error) {
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
