package entquery

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/gen/ent/joke"
	"github.com/abc-valera/netsly-api-golang/gen/ent/like"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
)

type likeQuery struct {
	*ent.Client
}

func NewLikeQuery(client *ent.Client) query.ILike {
	return &likeQuery{
		Client: client,
	}
}

func (lq *likeQuery) CountByJokeID(ctx context.Context, jokeID string) (int, error) {
	count, err := lq.Like.Query().
		Where(like.HasJokeWith(joke.ID(jokeID))).
		Count(ctx)
	return count, errhandler.HandleErr(err)
}
