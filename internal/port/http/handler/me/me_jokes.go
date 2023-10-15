package me

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/handler/other"
)

type MeJokesHandler struct {
	jokeRepo    repository.IJokeRepository
	jokeUseCase application.JokeUseCase
}

func NewMeJokesHandler(
	jokeRepo repository.IJokeRepository,
	jokeUseCase application.JokeUseCase,
) MeJokesHandler {
	return MeJokesHandler{
		jokeRepo:    jokeRepo,
		jokeUseCase: jokeUseCase,
	}
}

func (h MeJokesHandler) MeJokesGet(ctx context.Context, ogenParams ogen.MeJokesGetParams) (*ogen.Jokes, error) {
	userID := ctx.Value(other.PayloadKey).(service.Payload).UserID
	params, err := dto.NewDomainSelectParams(&ogenParams.SelectParams)
	if err != nil {
		return nil, err
	}
	domainJokes, err := h.jokeUseCase.GetJokesByUser(ctx, userID, params)
	return dto.NewJokesResponse(domainJokes), err
}

func (h MeJokesHandler) MeJokesPost(ctx context.Context, req *ogen.MeJokesPostReq) error {
	return h.jokeUseCase.CreateJoke(ctx, application.CreateJokeRequest{
		UserID:      ctx.Value(other.PayloadKey).(service.Payload).UserID,
		Title:       req.Title,
		Text:        req.Text,
		Explanation: req.Explanation.Value,
	})
}

func (h MeJokesHandler) MeJokesPut(ctx context.Context, req *ogen.MeJokesPutReq) error {
	return h.jokeUseCase.UpdateJoke(ctx, application.UpdateJokeRequest{
		JokeID:      req.JokeID,
		Title:       req.Title.Value,
		Text:        req.Text.Value,
		Explanation: req.Explanation.Value,
		UpdaterID:   ctx.Value(other.PayloadKey).(service.Payload).UserID,
	})
}

func (h MeJokesHandler) MeJokesDel(ctx context.Context, req *ogen.MeJokesDelReq) error {
	return h.jokeUseCase.DeleteJoke(ctx, application.DeleteJokeRequest{
		JokeID:    req.JokeID,
		DeleterID: ctx.Value(other.PayloadKey).(service.Payload).UserID,
	})
}
