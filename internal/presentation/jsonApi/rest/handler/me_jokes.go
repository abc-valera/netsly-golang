package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/contexts"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/restDto"
)

type MeJokesHandler struct {
	joke entity.IJoke
}

func NewMeJokesHandler(
	joke entity.IJoke,
) MeJokesHandler {
	return MeJokesHandler{
		joke: joke,
	}
}

func (h MeJokesHandler) MeJokesGet(ctx context.Context, ogenParams ogen.MeJokesGetParams) (*ogen.Jokes, error) {
	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	domainJokes, err := h.joke.GetAllByUserID(
		ctx,
		userID,
		restDto.NewDomainSelector(&ogenParams.Selector),
	)
	return restDto.NewJokes(domainJokes), err
}

func (h MeJokesHandler) MeJokesPost(ctx context.Context, req *ogen.MeJokesPostReq) (*ogen.Joke, error) {
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

func (h MeJokesHandler) MeJokesPut(ctx context.Context, req *ogen.MeJokesPutReq) (*ogen.Joke, error) {
	joke, err := h.joke.Update(ctx, req.JokeID, entity.JokeUpdateRequest{
		Title:       restDto.NewDomainOptionalString(req.Title),
		Text:        restDto.NewDomainOptionalString(req.Text),
		Explanation: restDto.NewDomainOptionalString(req.Explanation),
	})
	if err != nil {
		return nil, err
	}
	return restDto.NewJoke(joke), err
}

func (h MeJokesHandler) MeJokesDel(ctx context.Context, req *ogen.MeJokesDelReq) error {
	return h.joke.Delete(ctx, req.JokeID)
}
