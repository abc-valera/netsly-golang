package persistence

import (
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/impl"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
)

func NewRepositories(dbHost, dbPort, dbUser, dbPassword, dbName string) (repository.Repositories, error) {
	entRepos, err := impl.NewEntRepos(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		return repository.Repositories{}, err
	}

	return repository.NewRepositories(
		entRepos.UserRepository,
		entRepos.JokeRepository,
	)
}
