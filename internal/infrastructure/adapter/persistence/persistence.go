package persistence

import (
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/adapter/persistence/ent/impl"
)

func NewRepositories(dbHost, dbPort, dbUser, dbPassword, dbName string) (repository.Repositories, error) {
	userRepo, err := impl.NewEntImplementation(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		return repository.Repositories{}, err
	}

	return repository.Repositories{
		UserRepo: userRepo,
	}, nil
}
