package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/ent-impl/dto"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/persistence/ent-impl/errors"
)

type jokeCommand struct {
	*ent.Client
}

func NewJokeCommand(client *ent.Client) command.IJoke {
	return &jokeCommand{
		Client: client,
	}
}

func (jc jokeCommand) Create(ctx context.Context, req model.Joke) (model.Joke, error) {
	joke, err := jc.Joke.Create().
		SetID(req.ID).
		SetUserID(req.UserID).
		SetTitle(req.Title).
		SetText(req.Text).
		SetExplanation(req.Explanation).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return dto.FromEntJoke(joke), errhandler.HandleErr(err)
}

func (jc jokeCommand) Update(ctx context.Context, id string, req command.JokeUpdate) (model.Joke, error) {
	query := jc.Joke.UpdateOneID(id)
	if req.Title != nil {
		query.SetTitle(*req.Title)
	}
	if req.Text != nil {
		query.SetText(*req.Text)
	}
	if req.Explanation != nil {
		query.SetExplanation(*req.Explanation)
	}

	joke, err := query.
		Save(ctx)
	return dto.FromEntJoke(joke), errhandler.HandleErr(err)
}

func (jc jokeCommand) Delete(ctx context.Context, id string) error {
	err := jc.Joke.
		DeleteOneID(id).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
