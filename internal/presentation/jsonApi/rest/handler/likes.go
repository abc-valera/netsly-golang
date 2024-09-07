package handler

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"go.opentelemetry.io/otel/trace"
)

type LikesHandler struct {
	like entity.ILike
}

func NewLikesHandler(
	like entity.ILike,
) LikesHandler {
	return LikesHandler{
		like: like,
	}
}

func (h LikesHandler) LikesByJokeIDGet(ctx context.Context, params ogen.LikesByJokeIDGetParams) (int, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	return h.like.CountByJokeID(ctx, params.JokeID)
}
