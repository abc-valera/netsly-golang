package query

import "context"

type ILikeQuery interface {
	CountByJokeID(ctx context.Context, jokeID string) (int, error)
}
