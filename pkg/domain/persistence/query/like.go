package query

import "context"

type ILike interface {
	CountByJokeID(ctx context.Context, jokeID string) (int, error)
}
