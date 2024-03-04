package entcommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
	"github.com/abc-valera/netsly-api-golang/internal/persistence/ent-impl/dto"
	errhandler "github.com/abc-valera/netsly-api-golang/internal/persistence/ent-impl/errors"
)

type commentCommand struct {
	*ent.Client
}

func NewCommentCommand(client *ent.Client) command.IComment {
	return &commentCommand{
		Client: client,
	}
}

func (cc commentCommand) Create(ctx context.Context, req model.Comment) (model.Comment, error) {
	comment, err := cc.Comment.Create().
		SetID(req.ID).
		SetUserID(req.UserID).
		SetJokeID(req.JokeID).
		SetText(req.Text).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return dto.FromEntComment(comment), errhandler.HandleErr(err)
}

func (cc commentCommand) Update(ctx context.Context, id string, req command.CommentUpdate) (model.Comment, error) {
	query := cc.Comment.UpdateOneID(id)
	if req.Text != nil {
		query.SetText(*req.Text)
	}

	comment, err := query.
		Save(ctx)
	return dto.FromEntComment(comment), errhandler.HandleErr(err)
}

func (cc commentCommand) Delete(ctx context.Context, id string) error {
	err := cc.Comment.
		DeleteOneID(id).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
