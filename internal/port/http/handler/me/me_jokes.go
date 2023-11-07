package me

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/other"
)

type MeJokesHandler struct {
	jokeRepo repository.IJokeRepository
}

func NewMeJokesHandler(
	jokeRepo repository.IJokeRepository,
) MeJokesHandler {
	return MeJokesHandler{
		jokeRepo: jokeRepo,
	}
}

func (h MeJokesHandler) MeJokesGet(ctx context.Context, ogenParams ogen.MeJokesGetParams) (*ogen.Jokes, error) {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	params, err := dto.NewDomainSelectParams(&ogenParams.SelectParams)
	if err != nil {
		return nil, err
	}
	if err := repository.ValidateJokeSelectParams(params); err != nil {
		return nil, err
	}
	domainJokes, err := h.jokeRepo.GetByUserID(ctx, userID, params)
	return dto.NewJokesResponse(domainJokes), err
}

func (h MeJokesHandler) MeJokesPost(ctx context.Context, req *ogen.MeJokesPostReq) error {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	joke, err := entity.NewJoke(userID, req.Title, req.Text, req.Explanation.Value)
	if err != nil {
		return err
	}
	return h.jokeRepo.Create(ctx, joke)
}

func (h MeJokesHandler) MeJokesPut(ctx context.Context, req *ogen.MeJokesPutReq) error {
	updateReq, err := repository.NewJokeUpdateRequest(req.Title.Value, req.Text.Value, req.Explanation.Value)
	if err != nil {
		return err
	}
	return h.jokeRepo.Update(ctx, req.JokeID, updateReq)
}

func (h MeJokesHandler) MeJokesDel(ctx context.Context, req *ogen.MeJokesDelReq) error {
	return h.jokeRepo.Delete(ctx, req.JokeID)
}
