package repository

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/spec"
)

var (
	ErrJokeNotFound               = codeerr.NewMsgErr(codeerr.CodeNotFound, "Joke not found")
	ErrJokeWithTitleAlreadyExists = codeerr.NewMsgErr(codeerr.CodeAlreadyExists, "Joke with such title already exists for this user")

	ErrJokesOrderByNotSupported = codeerr.NewMsgErr(codeerr.CodeInvalidArgument, "OrderBy is supported only for 'title' and 'created_at' field")
)

type IJokeRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Joke, error)
	GetByUserID(ctx context.Context, userID string, spec spec.SelectParams) (entity.Jokes, error)
	Create(ctx context.Context, joke *entity.Joke) error
	Update(ctx context.Context, jokeID string, req JokeUpdateRequest) error
	Delete(ctx context.Context, jokeID string) error

	common.Transactioneer
}

func ValidateJokeSelectParams(params spec.SelectParams) error {
	if params.OrderBy != "" && params.OrderBy != "title" && params.OrderBy != "created_at" {
		return ErrJokesOrderByNotSupported
	}
	return nil
}

type JokeUpdateRequest struct {
	Title       string
	Text        string
	Explanation string
}

func NewJokeUpdateRequest(title, text, explanation string) (JokeUpdateRequest, error) {
	if title == "" {
		return JokeUpdateRequest{}, entity.ErrJokeTitleInvalid
	}
	if text == "" {
		return JokeUpdateRequest{}, entity.ErrJokeTextInvalid
	}

	return JokeUpdateRequest{
		Title:       title,
		Text:        text,
		Explanation: explanation,
	}, nil
}
