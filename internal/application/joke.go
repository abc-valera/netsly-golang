package application

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
)

var (
	errJokeModifyPermissionDenied = codeerr.NewMsgErr(codeerr.CodePermissionDenied, "You can modify only your own jokes")
)

type JokeUseCase struct {
	jokeRepo repository.JokeRepository
}

func NewJokeUseCase(jokeRepo repository.JokeRepository) JokeUseCase {
	return JokeUseCase{
		jokeRepo: jokeRepo,
	}
}

type UpdateJokeRequest struct {
	Joke   *entity.Joke
	UserID string
}

func (uc JokeUseCase) UpdateJoke(ctx context.Context, req UpdateJokeRequest) error {
	if req.Joke.UserID != req.UserID {
		return errJokeModifyPermissionDenied
	}

	return uc.jokeRepo.Update(ctx, req.Joke)
}

type DeleteJokeRequest struct {
	JokeID string
	UserID string
}

func (uc JokeUseCase) DeleteJoke(ctx context.Context, req DeleteJokeRequest) error {
	dbJoke, err := uc.jokeRepo.GetByID(ctx, req.JokeID)
	if err != nil {
		return err
	}

	if dbJoke.UserID != req.UserID {
		return errJokeModifyPermissionDenied
	}

	return uc.jokeRepo.Delete(ctx, req.JokeID)
}
