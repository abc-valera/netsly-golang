package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
)

var (
	ErrJokeWithIDNotFound = codeerr.NewMsgErr(codeerr.CodeNotFound, "Joke with such id not found")
)

type JokeRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Joke, error)
	GetByUserID(ctx context.Context, userID string) (entity.Jokes, error)
	Create(ctx context.Context, joke *entity.Joke) error
	Update(ctx context.Context, joke *entity.Joke) error
	Delete(ctx context.Context, id string) error

	Transactioneer
}
