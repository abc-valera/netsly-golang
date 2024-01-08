package main

import (
	"log"
	"os"

	"github.com/abc-valera/flugo-api-golang/internal/adapter/config"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/service"
	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	server "github.com/abc-valera/flugo-api-golang/internal/port/grpc"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = ".dev.env"
	}
	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal("Initialize config error: ", err)
	}

	commands, queries, tx, err := persistence.NewCommandsQueriesTx(config.DatabaseURL)
	if err != nil {
		log.Fatal("Initialize ent error: ", err)
	}

	services, err := service.NewServices(
		config.AccessTokenDuration, config.RefreshTokenDuration,
		config.RedisPort, config.RedisUser, config.RedisPass,
	)
	if err != nil {
		log.Fatal("Initialize services error: ", err)
	}

	domains := domain.NewDomains(commands, queries, services)

	usecases, err := application.NewUseCases(queries, tx, domains, services)
	if err != nil {
		log.Fatal("Initialize usecases error: ", err)
	}

	if err := server.RunServer(config.GRPCPort, services, usecases); err != nil {
		log.Fatal("gRPC server error: ", err)
	}
}
