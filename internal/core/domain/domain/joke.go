package domain

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/coderr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
)

var (
	ErrJokeIDInvalid          = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid joke ID")
	ErrJokeUserIDInvalid      = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid user ID for joke")
	ErrJokeTitleInvalid       = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid title")
	ErrJokeTextInvalid        = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid text")
	ErrJokeExplanationInvalid = coderr.NewMessage(coderr.CodeInvalidArgument, "Provided invalid explanation")
)

type Joke struct {
	command command.IJoke
}

func NewJoke(
	command command.IJoke,
) Joke {
	return Joke{
		command: command,
	}
}

type JokeCreateRequest struct {
	UserID      string
	Title       string
	Text        string
	Explanation string
}

func (j Joke) Create(ctx context.Context, req JokeCreateRequest) error {
	// Validation
	if req.UserID == "" {
		return ErrCommentUserIDInvalid
	}
	if req.Title == "" || len(req.Title) < 4 || len(req.Title) > 64 {
		return ErrJokeTitleInvalid
	}
	if req.Text == "" || len(req.Text) < 4 || len(req.Text) > 4096 {
		return ErrJokeTextInvalid
	}
	if len(req.Explanation) > 4096 {
		return ErrJokeExplanationInvalid
	}

	// Domain logic
	baseModel := common.NewBaseModel()

	return j.command.Create(ctx, model.Joke{
		BaseModel:   baseModel,
		UserID:      req.UserID,
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
	})
}

type JokeUpdateRequest struct {
	Title       *string
	Text        *string
	Explanation *string
}

func (j Joke) Update(ctx context.Context, jokeID string, req JokeUpdateRequest) error {
	// Validation
	if jokeID == "" {
		return ErrJokeIDInvalid
	}
	if req.Title != nil || len(*req.Title) < 4 || len(*req.Title) > 64 {
		return ErrJokeTitleInvalid
	}
	if req.Text != nil || len(*req.Text) < 4 || len(*req.Text) > 4096 {
		return ErrJokeTextInvalid
	}
	if req.Explanation != nil || len(*req.Explanation) > 4096 {
		return ErrJokeExplanationInvalid
	}

	// Edit in data source
	return j.command.Update(ctx, jokeID, command.JokeUpdate{
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation,
	})
}

func (j Joke) Delete(ctx context.Context, jokeID string) error {
	// Validation
	if jokeID == "" {
		return ErrJokeIDInvalid
	}

	// Delete in data source
	return j.command.Delete(ctx, jokeID)
}
