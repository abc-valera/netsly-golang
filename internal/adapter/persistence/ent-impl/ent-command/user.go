package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/dto"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

type userCommand struct {
	*ent.Client
}

func NewUserCommand(client *ent.Client) command.IUser {
	return &userCommand{
		Client: client,
	}
}

func (uc userCommand) Create(ctx context.Context, req model.User) (model.User, error) {
	user, err := uc.User.Create().
		SetID(req.ID).
		SetUsername(req.Username).
		SetEmail(req.Email).
		SetHashedPassword(req.HashedPassword).
		SetFullname(req.Fullname).
		SetStatus(req.Status).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return dto.FromEntUser(user), errhandler.HandleErr(err)
}

func (uc userCommand) Update(ctx context.Context, id string, req command.UserUpdate) (model.User, error) {
	query := uc.User.UpdateOneID(id)
	if req.HashedPassword != nil {
		query.SetHashedPassword(*req.HashedPassword)
	}
	if req.Fullname != nil {
		query.SetFullname(*req.Fullname)
	}
	if req.Status != nil {
		query.SetStatus(*req.Status)
	}

	user, err := query.
		Save(ctx)
	return dto.FromEntUser(user), errhandler.HandleErr(err)
}

func (uc userCommand) Delete(ctx context.Context, id string) error {
	err := uc.User.
		DeleteOneID(id).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
