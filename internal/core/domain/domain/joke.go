package domain

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
)

var (
	ErrJokeIDInvalid          = codeerr.NewMessage(codeerr.CodeInvalidArgument, "Provided invalid joke ID")
	ErrJokeUserIDInvalid      = codeerr.NewMessage(codeerr.CodeInvalidArgument, "Provided invalid user ID for joke")
	ErrJokeTitleInvalid       = codeerr.NewMessage(codeerr.CodeInvalidArgument, "Provided invalid title")
	ErrJokeTextInvalid        = codeerr.NewMessage(codeerr.CodeInvalidArgument, "Provided invalid text")
	ErrJokeExplanationInvalid = codeerr.NewMessage(codeerr.CodeInvalidArgument, "Provided invalid explanation")
)

type JokeDomain struct {
	command command.IJokeCommand
}

func NewJokeDomain(
	command command.IJokeCommand,
) JokeDomain {
	return JokeDomain{
		command: command,
	}
}

type JokeCreateRequest struct {
	UserID      string
	Title       string
	Text        string
	Explanation string
}

func (j JokeDomain) Create(ctx context.Context, req JokeCreateRequest) error {
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

func (j JokeDomain) Update(ctx context.Context, jokeID string, req JokeUpdateRequest) error {
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

func (j JokeDomain) Delete(ctx context.Context, jokeID string) error {
	// Validation
	if jokeID == "" {
		return ErrJokeIDInvalid
	}

	// Delete in data source
	return j.command.Delete(ctx, jokeID)
}
