package application

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/spec"
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

type CreateJokeRequest struct {
	UserID      string
	Title       string
	Text        string
	Explanation string
}

func (uc JokeUseCase) CreateJoke(ctx context.Context, req CreateJokeRequest) error {
	domainJoke, err := entity.NewJoke(req.UserID, req.Title, req.Text, req.Explanation)
	if err != nil {
		return err
	}
	return uc.jokeRepo.Create(ctx, domainJoke)
}

func (uc JokeUseCase) GetJokesByUser(ctx context.Context, userID string, params spec.SelectParams) (entity.Jokes, error) {
	if err := entity.ValidateJokeSelectParams(params); err != nil {
		return nil, err
	}
	return uc.jokeRepo.GetByUserID(ctx, userID, params)
}

type UpdateJokeRequest struct {
	JokeID      string
	Title       string
	Text        string
	Explanation string
	UpdaterID   string
}

func (uc JokeUseCase) UpdateJoke(ctx context.Context, req UpdateJokeRequest) error {
	domainJoke, err := uc.jokeRepo.GetByID(ctx, req.JokeID)
	if err != nil {
		return err
	}

	if req.UpdaterID != domainJoke.UserID {
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
	JokeID    string
	DeleterID string
}

func (uc JokeUseCase) DeleteJoke(ctx context.Context, req DeleteJokeRequest) error {
	dbJoke, err := uc.jokeRepo.GetByID(ctx, req.JokeID)
	if err != nil {
		return err
	}

	if dbJoke.UserID != req.DeleterID {
		return errJokeModifyPermissionDenied
	}

	return uc.jokeRepo.Delete(ctx, req.JokeID)
}
