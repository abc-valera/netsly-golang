package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/common"
)

var (
	ErrJokeNotFound               = codeerr.NewMsgErr(codeerr.CodeNotFound, "Joke not found")
	ErrJokeWithTitleAlreadyExists = codeerr.NewMsgErr(codeerr.CodeAlreadyExists, "Joke with such title already exists for this user")
)

type JokeRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Joke, error)
	GetByUserID(ctx context.Context, userID string) (entity.Jokes, error)
	Create(ctx context.Context, joke *entity.Joke) error
	Update(ctx context.Context, joke *entity.Joke) error
	Delete(ctx context.Context, jokeID string) error

	common.Transactioneer
}
