package command

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/model"
)

type IUserCommand interface {
	Create(ctx context.Context, req model.User) error
	Update(ctx context.Context, id string, req model.UserUpdate) error
	Delete(ctx context.Context, id string) error
}
