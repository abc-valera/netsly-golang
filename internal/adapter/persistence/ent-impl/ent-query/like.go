package entquery

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/like"
	errhandler "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
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
		Where(like.JokeID(jokeID)).
		Count(ctx)
	return count, errhandler.HandleErr(err)
}
