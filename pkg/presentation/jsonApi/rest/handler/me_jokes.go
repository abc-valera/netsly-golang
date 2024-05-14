package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/rest/restDto"
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
	domainJokes, err := h.joke.GetAllByUserID(
		ctx,
		payloadUserID(ctx),
		restDto.NewDomainSelector(&ogenParams.Selector),
	)
	return restDto.NewJokesResponse(domainJokes), err
}

func (h MeJokesHandler) MeJokesPost(ctx context.Context, req *ogen.MeJokesPostReq) (*ogen.Joke, error) {
	joke, err := h.joke.Create(ctx, entity.JokeCreateRequest{
		UserID:      payloadUserID(ctx),
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation.Value,
	})
	if err != nil {
		return nil, err
	}
	return restDto.NewJokeResponse(joke), err
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
	return restDto.NewJokeResponse(joke), err
}

func (h MeJokesHandler) MeJokesDel(ctx context.Context, req *ogen.MeJokesDelReq) error {
	return h.joke.Delete(ctx, req.JokeID)
}
