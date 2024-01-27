package entcommand

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/joke"
	errhandler "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/persistence/command"
)

type jokeCommand struct {
	*ent.Client
}

func NewJokeCommand(client *ent.Client) command.IJoke {
	return &jokeCommand{
		Client: client,
	}
}

func (jc jokeCommand) Create(ctx context.Context, req model.Joke) error {
	_, err := jc.Joke.Create().
		SetID(req.ID).
		SetUserID(req.UserID).
		SetText(req.Text).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return errhandler.HandleErr(err)
}

func (jc jokeCommand) Update(ctx context.Context, id string, req command.JokeUpdate) error {
	query := jc.Joke.Update()
	if req.Title != nil {
		query.SetTitle(*req.Title)
	}
	if req.Text != nil {
		query.SetText(*req.Text)
	}
	if req.Explanation != nil {
		query.SetExplanation(*req.Explanation)
	}

	_, err := query.
		Where(joke.ID(id)).
		Save(ctx)
	return errhandler.HandleErr(err)
}

func (jc jokeCommand) Delete(ctx context.Context, id string) error {
	err := jc.Joke.
		DeleteOneID(id).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
