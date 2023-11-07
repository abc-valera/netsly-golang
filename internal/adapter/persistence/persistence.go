package persistence

import (
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/impl"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
)

func NewRepositories(databaseURL string) (repository.Repositories, error) {
	entRepos, err := impl.NewEntRepos(databaseURL)
	if err != nil {
		return repository.Repositories{}, err
	}

	return repository.NewRepositories(
		entRepos.IUserRepository,
		entRepos.IJokeRepository,
		entRepos.ICommentRepository,
		entRepos.ILikeRepository,
	)
}
