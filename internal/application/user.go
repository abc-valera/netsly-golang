package application

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
)

// Rewrite the UseCase from internal/application/joke.go for user:

var (
	errUserModifyPermissionDenied = codeerr.NewMsgErr(codeerr.CodePermissionDenied, "You can modify only your own user")
)

type UserUseCase struct {
	userRepo repository.IUserRepository
}

func NewUserUseCase(userRepo repository.IUserRepository) UserUseCase {
	return UserUseCase{
		userRepo: userRepo,
	}
}

type UpdateUserRequest struct {
	UpdaterID string
	UserID    string
	Username  string
	Fullanme  string
	Status    string
}

func (uc UserUseCase) UpdateUser(ctx context.Context, req UpdateUserRequest) error {
	domainUser, err := uc.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if req.UpdaterID != domainUser.ID {
		return errUserModifyPermissionDenied
	}

	if req.Username != "" {
		domainUser.Username = req.Username
	}
	if req.Fullanme != "" {
		domainUser.Fullname = req.Fullanme
	}
	if req.Status != "" {
		domainUser.Status = req.Status
	}

	return uc.userRepo.Update(ctx, domainUser)
}

type DeleteUserRequest struct {
	DeleterID string
	UserID    string
}

func (uc UserUseCase) DeleteUser(ctx context.Context, req DeleteUserRequest) error {
	domainUser, err := uc.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if domainUser.ID != req.DeleterID {
		return errUserModifyPermissionDenied
	}

	return uc.userRepo.Delete(ctx, req.UserID)
}
