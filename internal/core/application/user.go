package application

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
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
	UserID   string
	Username string
	Fullanme string
	Status   string
}

func (uc UserUseCase) UpdateUser(ctx context.Context, req UpdateUserRequest) error {
	domainUser, err := uc.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return err
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
