package impl

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/user"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/impl/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
)

type userRepository struct {
	common.BaseRepository
}

func NewUserRepository(client *ent.Client) repository.IUserRepository {
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

func (r userRepository) Update(ctx context.Context, userID string, req repository.UserUpdateRequest) error {
	query := r.Client.User.Update()
	if req.Username != "" {
		query.SetUsername(req.Username)
	}
	if req.Email != "" {
		query.SetEmail(req.Email)
	}
	if req.Password != "" {
		query.SetHashedPassword(req.Password)
	}
	if req.Fullname != "" {
		query.SetFullname(req.Fullname)
	}
	if req.Status != "" {
		query.SetStatus(req.Status)
	}

	_, err := query.
		Where(user.ID(userID)).
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
