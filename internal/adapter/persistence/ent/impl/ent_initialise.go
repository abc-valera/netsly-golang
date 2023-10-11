package impl

import (
	"context"
	"fmt"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"

	_ "github.com/lib/pq"
)

func NewEntRepos(dbHost, dbPort, dbUser, dbPassword, dbName string) (struct {
	repository.IUserRepository
	repository.IJokeRepository
	// repository.CommentRepository
	// repository.LikeRepository
}, error) {
	repos := struct {
		repository.IUserRepository
		repository.IJokeRepository
		// repository.CommentRepository
		// repository.LikeRepository
	}{}

	// Connect to the database
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword))
	if err != nil {
		return repos, codeerr.NewInternal("NewEntImplementation", err)
	}

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		return repos, codeerr.NewInternal("NewEntImplementation", err)
	}

	repos.IUserRepository = NewUserRepository(client)
	repos.IJokeRepository = NewJokeRepository(client)
	// repos.CommentRepository = NewCommentRepository(client)
	// repos.LikeRepository = NewLikeRepository(client)
	return repos, nil
}
