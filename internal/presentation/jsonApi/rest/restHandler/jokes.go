package restHandler

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/selector"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/contexts"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/restDto"
	"go.opentelemetry.io/otel/trace"
)

type Jokes struct {
	joke entity.IJoke
}

func newJokes(
	joke entity.IJoke,
) Jokes {
	return Jokes{
		joke: joke,
	}
}

func (h Jokes) JokesGet(ctx context.Context, ogenParams ogen.JokesGetParams) (ogen.Jokes, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	jokes, err := h.joke.GetMany(
		ctx,
		selector.WithFilter(model.Joke{UserID: userID}),
		selector.WithLimit[model.Joke](uint(*restDto.NewInt(ogenParams.Limit))),
		selector.WithOffset[model.Joke](uint(*restDto.NewInt(ogenParams.Offset))),
	)

	return restDto.NewJokes(jokes), err
}

func (h Jokes) JokesPost(ctx context.Context, req *ogen.JokesPostReq) (*ogen.Joke, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	joke, err := h.joke.Create(ctx, entity.JokeCreateRequest{
		UserID:      userID,
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation.Value,
	})
	if err != nil {
		return nil, err
	}
	return restDto.NewJoke(joke), err
}

func (h Jokes) JokesPut(ctx context.Context, req *ogen.JokesPutReq) (*ogen.Joke, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	joke, err := h.joke.Update(ctx, req.JokeID, entity.JokeUpdateRequest{
		Title:       restDto.NewString(req.Title),
		Text:        restDto.NewString(req.Text),
		Explanation: restDto.NewString(req.Explanation),
	})
	if err != nil {
		return nil, err
	}
	return restDto.NewJoke(joke), err
}

func (h Jokes) JokesDel(ctx context.Context, req *ogen.JokesDelReq) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	return h.joke.Delete(ctx, req.JokeID)
}
