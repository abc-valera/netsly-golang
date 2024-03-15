package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/dto"
)

type MeJokesHandler struct {
	jokeQuery  query.IJoke
	jokeEntity entity.IJoke
}

func NewMeJokesHandler(
	jokeQuery query.IJoke,
	jokeEntity entity.IJoke,
) MeJokesHandler {
	return MeJokesHandler{
		jokeQuery:  jokeQuery,
		jokeEntity: jokeEntity,
	}
}

func (h MeJokesHandler) MeJokesGet(ctx context.Context, ogenParams ogen.MeJokesGetParams) (*ogen.Jokes, error) {
	domainJokes, err := h.jokeQuery.GetAllByUserID(
		ctx,
		payloadUserID(ctx),
		dto.NewDomainSelectParams(&ogenParams.SelectParams),
	)
	return dto.NewJokesResponse(domainJokes), err
}

func (h MeJokesHandler) MeJokesPost(ctx context.Context, req *ogen.MeJokesPostReq) (*ogen.Joke, error) {
	joke, err := h.jokeEntity.Create(ctx, entity.JokeCreateRequest{
		UserID:      payloadUserID(ctx),
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation.Value,
	})
	if err != nil {
		return nil, err
	}
	return dto.NewJokeResponse(joke), err
}

func (h MeJokesHandler) MeJokesPut(ctx context.Context, req *ogen.MeJokesPutReq) (*ogen.Joke, error) {
	joke, err := h.jokeEntity.Update(ctx, req.JokeID, entity.JokeUpdateRequest{
		Title:       dto.NewPointerString(req.Title),
		Text:        dto.NewPointerString(req.Text),
		Explanation: dto.NewPointerString(req.Explanation),
	})
	if err != nil {
		return nil, err
	}
	return dto.NewJokeResponse(joke), err
}

func (h MeJokesHandler) MeJokesDel(ctx context.Context, req *ogen.MeJokesDelReq) error {
	return h.jokeEntity.Delete(ctx, req.JokeID)
}
