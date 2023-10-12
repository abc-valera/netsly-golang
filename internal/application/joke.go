package application

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
)

var (
	errJokeModifyPermissionDenied = codeerr.NewMsgErr(codeerr.CodePermissionDenied, "You can modify only your own jokes")
)

type JokeUseCase struct {
	jokeRepo repository.IJokeRepository
}

func NewJokeUseCase(jokeRepo repository.IJokeRepository) JokeUseCase {
	return JokeUseCase{
		jokeRepo: jokeRepo,
	}
}

type UpdateJokeRequest struct {
	JokeID      string
	Title       string
	Text        string
	Explanation string
	UserID      string
}

func (uc JokeUseCase) UpdateJoke(ctx context.Context, req UpdateJokeRequest) error {
	domainJoke, err := uc.jokeRepo.GetByID(ctx, req.JokeID)
	if err != nil {
		return err
	}

	if req.UserID != domainJoke.UserID {
		return errJokeModifyPermissionDenied
	}

	if req.Title != "" {
		domainJoke.Title = req.Title
	}
	if req.Text != "" {
		domainJoke.Text = req.Text
	}
	if req.Explanation != "" {
		domainJoke.Explanation = req.Explanation
	}

	return uc.jokeRepo.Update(ctx, domainJoke)
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
