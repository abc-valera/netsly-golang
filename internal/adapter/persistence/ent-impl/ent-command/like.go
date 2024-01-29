package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/like"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model"
)

type likeCommand struct {
	*ent.Client
}

func NewLikeCommand(client *ent.Client) command.ILike {
	return &likeCommand{
		Client: client,
	}
}

func (lc likeCommand) Create(ctx context.Context, req model.Like) error {
	_, err := lc.Like.Create().
		SetUserID(req.UserID).
		SetJokeID(req.JokeID).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return errhandler.HandleErr(err)
}

func (lc likeCommand) Delete(ctx context.Context, userID string, jokeID string) error {
	_, err := lc.Like.Delete().
		Where(like.UserID(userID), like.JokeID(jokeID)).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
