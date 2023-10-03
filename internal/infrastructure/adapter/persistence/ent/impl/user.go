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
	client *ent.Client
}

func NewUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
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

func (r *userRepository) PerformTX(ctx context.Context, txFunc func(ctx context.Context) error) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return handleErr(err)
	}

	oldClient := r.client
	defer func() {
		r.client = oldClient
	}()

	r.client = tx.Client()
	if err := txFunc(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return handleErr(err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return handleErr(err)
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	entUser, err := r.client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	return dto.FromEntUserToUser(entUser), handleErr(err)
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	entUser, err := r.client.User.
		Query().
		Where(user.Username(username)).
		Only(ctx)
	return dto.FromEntUserToUser(entUser), handleErr(err)
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	entUser, err := r.client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	return dto.FromEntUserToUser(entUser), handleErr(err)
}

func (r *userRepository) Update(ctx context.Context, domainUser *entity.User) error {
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

func (r *userRepository) Delete(ctx context.Context, id string) error {
	return handleErr(
		r.client.User.
			DeleteOneID(id).
			Exec(ctx),
	)
}
