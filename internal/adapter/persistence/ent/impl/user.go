package impl

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/user"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/impl/common"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
)

type userRepository struct {
	common.BaseRepository
}

func NewUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{
		BaseRepository: common.NewBaseRepository(client),
	}
}

func (r userRepository) Create(ctx context.Context, user *entity.User) error {
	_, err := r.Client.User.
		Create().
		SetID(user.ID).
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetHashedPassword(user.HashedPassword).
		SetFullname(user.Fullname).
		SetStatus(user.Status).
		SetCreatedAt(user.CreatedAt).
		Save(ctx)
	return common.HandleErr(err)
}

func (r userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	entUser, err := r.Client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	return dto.FromEntUserToUser(entUser), common.HandleErr(err)
}

func (r userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	entUser, err := r.Client.User.
		Query().
		Where(user.Username(username)).
		Only(ctx)
	return dto.FromEntUserToUser(entUser), common.HandleErr(err)
}

func (r userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	entUser, err := r.Client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	return dto.FromEntUserToUser(entUser), common.HandleErr(err)
}

func (r userRepository) Update(ctx context.Context, domainUser *entity.User) error {
	_, err := r.Client.User.
		Update().
		Where(user.ID(domainUser.ID)).
		SetUsername(domainUser.Username).
		SetEmail(domainUser.Email).
		SetHashedPassword(domainUser.HashedPassword).
		SetFullname(domainUser.Fullname).
		SetStatus(domainUser.Status).
		Save(ctx)
	return common.HandleErr(err)
}

func (r userRepository) Delete(ctx context.Context, id string) error {
	return common.HandleErr(
		r.Client.User.
			DeleteOneID(id).
			Exec(ctx),
	)
}
