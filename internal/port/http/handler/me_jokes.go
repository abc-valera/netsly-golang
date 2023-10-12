package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/application"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
)

type MeJokesHandler struct {
	jokeRepo    repository.IJokeRepository
	jokeUsecase application.JokeUseCase
}

func NewMeJokesHandler(
	jokeRepo repository.IJokeRepository,
	jokeUsecase application.JokeUseCase,
) MeJokesHandler {
	return MeJokesHandler{
		jokeRepo:    jokeRepo,
		jokeUsecase: jokeUsecase,
	}
}

func (h MeJokesHandler) MeJokesGet(ctx context.Context) (*ogen.Jokes, error) {
	userID := ctx.Value(PayloadKey).(service.Payload).UserID
	jokes, err := h.jokeRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dto.NewJokesResponse(jokes), nil
}

func (h MeJokesHandler) MeJokesPost(ctx context.Context, req *ogen.MeJokesPostReq) error {
	userID := ctx.Value(PayloadKey).(service.Payload).UserID
	domainJoke, err := entity.NewJoke(userID, req.Title, req.Text, req.Explanation.Value)
	if err != nil {
		return err
	}
	return h.jokeRepo.Create(ctx, domainJoke)
}

func (h MeJokesHandler) MeJokesPut(ctx context.Context, req *ogen.MeJokesPutReq) error {
	userID := ctx.Value(PayloadKey).(service.Payload).UserID
	return h.jokeUsecase.UpdateJoke(ctx, application.UpdateJokeRequest{
		JokeID:      req.JokeID,
		Title:       req.Title.Value,
		Text:        req.Text.Value,
		Explanation: req.Explanation.Value,
		UserID:      userID,
	})
}

func (h MeJokesHandler) MeJokesDelete(ctx context.Context, req *ogen.MeJokesDeleteReq) error {
	userID := ctx.Value(PayloadKey).(service.Payload).UserID
	return h.jokeUsecase.DeleteJoke(ctx, application.DeleteJokeRequest{
		UserID: userID,
	})
}
