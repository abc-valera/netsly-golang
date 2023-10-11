package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
)

type MeJokesHandler struct {
	userRepo repository.UserRepository
	jokeRepo repository.JokeRepository
}

func NewMeJokesHandler(
	userRepo repository.UserRepository,
	jokeRepo repository.JokeRepository,
) MeJokesHandler {
	return MeJokesHandler{
		userRepo: userRepo,
		jokeRepo: jokeRepo,
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

// MeJokesPost(ctx context.Context, req *Joke) error
func (h MeJokesHandler) MeJokesPost(ctx context.Context, req *ogen.MeJokesPostReq) error {
	userID := ctx.Value(PayloadKey).(service.Payload).UserID
	domainJoke, err := entity.NewJoke(userID, req.Title, req.Text, req.Explanation.Value)
	if err != nil {
		return err
	}
	if err := h.jokeRepo.Create(ctx, domainJoke); err != nil {
		return err
	}
	return nil
}
