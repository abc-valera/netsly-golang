package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
)

type MeJokesHandler struct {
	jokeQuery  query.IJoke
	jokeDomain domain.Joke
}

func NewMeJokesHandler(
	jokeQuery query.IJoke,
	jokeDomain domain.Joke,
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
	return h.jokeDomain.Create(ctx, domain.JokeCreateRequest{
		UserID:      payloadUserID(ctx),
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation.Value,
	})
}

func (h MeJokesHandler) MeJokesPut(ctx context.Context, req *ogen.MeJokesPutReq) error {
	return h.jokeDomain.Update(ctx, req.JokeID, domain.JokeUpdateRequest{
		Title:       dto.NewPointerString(req.Title),
		Text:        dto.NewPointerString(req.Text),
		Explanation: dto.NewPointerString(req.Explanation),
	})
}

func (h MeJokesHandler) MeJokesDel(ctx context.Context, req *ogen.MeJokesDelReq) error {
	return h.jokeDomain.Delete(ctx, req.JokeID)
}
