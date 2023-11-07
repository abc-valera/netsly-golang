package main

import (
	"log"
	"os"

	"github.com/abc-valera/flugo-api-golang/internal/adapter/config"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service"
	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/port/http"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = ".dev.env"
	}
	config, err := config.InitConfig(configPath)
	if err != nil {
		log.Fatal("Initialize config error: ", err)
	}

	repos, err := persistence.NewRepositories(config.DatabaseURL)
	if err != nil {
		log.Fatal("Initialize postgre error: ", err)
	}

	services, err := service.NewServices(
		config.AccessTokenDuration, config.RefreshTokenDuration,
		config.RedisPort, config.RedisUser, config.RedisPass,
	)
	if err != nil {
		log.Fatal("Initialize services error: ", err)
	}

	usecases, err := application.NewUseCases(repos, services)
	if err != nil {
		log.Fatal("Initialize usecases error: ", err)
	}

	if err := http.RunServer(config.HTTPPort, config.HTTPDocsPath, repos, services, usecases); err != nil {
		log.Fatal("HTTP server error: ", err)
	}
}
