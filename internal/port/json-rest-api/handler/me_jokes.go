package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/dto"
)

type MeJokesHandler struct {
	jokeQuery  query.IJoke
	jokeDomain entity.Joke
}

func NewMeJokesHandler(
	jokeQuery query.IJoke,
	jokeDomain entity.Joke,
) MeJokesHandler {
	return MeJokesHandler{
		jokeQuery:  jokeQuery,
		jokeDomain: jokeDomain,
	}
}

func (h MeJokesHandler) MeJokesGet(ctx context.Context, ogenParams ogen.MeJokesGetParams) (*ogen.Jokes, error) {
	userID := payloadUserID(ctx)
	params, err := dto.NewDomainSelectParams(&ogenParams.SelectParams)
	if err != nil {
		return nil, err
	}
	domainJokes, err := h.jokeQuery.GetAllByUserID(ctx, userID, params)
	return dto.NewJokesResponse(domainJokes), err
}

func (h MeJokesHandler) MeJokesPost(ctx context.Context, req *ogen.MeJokesPostReq) error {
	return h.jokeDomain.Create(ctx, entity.JokeCreateRequest{
		UserID:      payloadUserID(ctx),
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation.Value,
	})
}

func (h MeJokesHandler) MeJokesPut(ctx context.Context, req *ogen.MeJokesPutReq) error {
	return h.jokeDomain.Update(ctx, req.JokeID, entity.JokeUpdateRequest{
		Title:       dto.NewPointerString(req.Title),
		Text:        dto.NewPointerString(req.Text),
		Explanation: dto.NewPointerString(req.Explanation),
	})
}

func (h MeJokesHandler) MeJokesDel(ctx context.Context, req *ogen.MeJokesDelReq) error {
	return h.jokeDomain.Delete(ctx, req.JokeID)
}
