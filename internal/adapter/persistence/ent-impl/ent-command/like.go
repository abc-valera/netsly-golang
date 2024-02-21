package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/like"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/dto"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type likeCommand struct {
	*ent.Client
}

func NewLikeCommand(client *ent.Client) command.ILike {
	return &likeCommand{
		Client: client,
	}
}

func (lc likeCommand) Create(ctx context.Context, req model.Like) (model.Like, error) {
	like, err := lc.Like.Create().
		SetUserID(req.UserID).
		SetJokeID(req.JokeID).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return dto.FromEntLike(like), errhandler.HandleErr(err)
}

func (lc likeCommand) Delete(ctx context.Context, userID string, jokeID string) error {
	_, err := lc.Like.Delete().
		Where(like.UserID(userID), like.JokeID(jokeID)).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
