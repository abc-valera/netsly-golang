package impl

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/user"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/infrastructure/adapter/persistence/ent/dto"
)

type userRepository struct {
	baseRepository
}

func NewUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{
		baseRepository: NewBaseRepository(client),
	}
}

func (r userRepository) Create(ctx context.Context, user *entity.User) error {
	_, err := r.client.User.
		Create().
		SetID(user.ID).
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetHashedPassword(user.HashedPassword).
		SetFullname(user.Fullname).
		SetStatus(user.Status).
		SetCreatedAt(user.CreatedAt).
		Save(ctx)
	return handleErr(err)
}

func (r userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	entUser, err := r.client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	return dto.FromEntUserToUser(entUser), handleErr(err)
}

func (r userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	entUser, err := r.client.User.
		Query().
		Where(user.Username(username)).
		Only(ctx)
	return dto.FromEntUserToUser(entUser), handleErr(err)
}

func (r userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	entUser, err := r.client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	return dto.FromEntUserToUser(entUser), handleErr(err)
}

func (r userRepository) Update(ctx context.Context, domainUser *entity.User) error {
	_, err := r.client.User.
		Update().
		Where(user.ID(domainUser.ID)).
		SetUsername(domainUser.Username).
		SetEmail(domainUser.Email).
		SetHashedPassword(domainUser.HashedPassword).
		SetFullname(domainUser.Fullname).
		SetStatus(domainUser.Status).
		Save(ctx)
	return handleErr(err)
}

func (r userRepository) Delete(ctx context.Context, id string) error {
	return handleErr(
		r.client.User.
			DeleteOneID(id).
			Exec(ctx),
	)
}
