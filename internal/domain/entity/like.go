package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"go.opentelemetry.io/otel/trace"
)

type ILike interface {
	Create(ctx context.Context, req LikeCreateRequest) (model.Like, error)
	Delete(ctx context.Context, userID, jokeID string) error

	query.ILike
}

type like struct {
	IDependency

	query.ILike
}

func newLike(dep IDependency) ILike {
	return like{
		IDependency: dep,

		ILike: dep.Q().Like,
	}
}

type LikeCreateRequest struct {
	UserID string `validate:"required,uuid"`
	JokeID string `validate:"required,uuid"`
}

func (e like) Create(ctx context.Context, req LikeCreateRequest) (model.Like, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.Like{}, err
	}

	like := model.Like{
		CreatedAt: time.Now().Truncate(time.Millisecond).Local(),
		UserID:    req.UserID,
		JokeID:    req.JokeID,
	}

	if err := e.C().Like.Create(ctx, like); err != nil {
		return model.Like{}, err
	}

	return like, nil
}

func (e like) Delete(ctx context.Context, userID string, jokeID string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	return e.C().Like.Delete(ctx, model.Like{UserID: userID, JokeID: jokeID})
}
