package command

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

var ErrLikeAlreadyExists = coderr.NewCodeMessage(coderr.CodeAlreadyExists, "Like already exists")

type ILike interface {
	Create(ctx context.Context, req model.Like) (model.Like, error)
	Delete(ctx context.Context, req model.Like) error
}
