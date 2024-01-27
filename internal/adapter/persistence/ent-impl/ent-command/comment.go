package entcommand

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/comment"
	errhandler "github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent-impl/errors"
	"github.com/abc-valera/flugo-api-golang/internal/core/model"
	"github.com/abc-valera/flugo-api-golang/internal/core/persistence/command"
)

type commentCommand struct {
	*ent.Client
}

func NewCommentCommand(client *ent.Client) command.IComment {
	return &commentCommand{
		Client: client,
	}
}

func (cc commentCommand) Create(ctx context.Context, req model.Comment) error {
	_, err := cc.Comment.Create().
		SetID(req.ID).
		SetUserID(req.UserID).
		SetJokeID(req.JokeID).
		SetText(req.Text).
		SetCreatedAt(req.CreatedAt).
		Save(ctx)
	return errhandler.HandleErr(err)
}

func (cc commentCommand) Update(ctx context.Context, commentID string, req command.CommentUpdate) error {
	query := cc.Comment.Update()
	if req.Text != nil {
		query.SetText(*req.Text)
	}

	_, err := query.
		Where(comment.ID(commentID)).
		Save(ctx)
	return errhandler.HandleErr(err)
}

func (cc commentCommand) Delete(ctx context.Context, id string) error {
	err := cc.Comment.
		DeleteOneID(id).
		Exec(ctx)
	return errhandler.HandleErr(err)
}
