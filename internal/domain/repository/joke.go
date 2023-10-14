package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/common"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/spec"
)

var (
	ErrJokeNotFound               = codeerr.NewMsgErr(codeerr.CodeNotFound, "Joke not found")
	ErrJokeWithTitleAlreadyExists = codeerr.NewMsgErr(codeerr.CodeAlreadyExists, "Joke with such title already exists for this user")
)

type IJokeRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Joke, error)
	GetByUserID(ctx context.Context, userID string, spec spec.SelectParams) (entity.Jokes, error)
	Create(ctx context.Context, joke *entity.Joke) error
	Update(ctx context.Context, joke *entity.Joke) error
	Delete(ctx context.Context, jokeID string) error

	common.Transactioneer
}
