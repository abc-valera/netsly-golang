package impl

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/like"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/impl/common"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
)

type likeRepository struct {
	common.BaseRepository
}

func NewLikeRepository(client *ent.Client) repository.ILikeRepository {
	return &likeRepository{
		BaseRepository: common.NewBaseRepository(client),
	}
}

func (r likeRepository) CountByJokeID(ctx context.Context, jokeID string) (int, error) {
	return r.Client.Like.
		Query().
		Where(like.JokeID(jokeID)).
		Count(ctx)
}

func (r likeRepository) Create(ctx context.Context, like *entity.Like) error {
	_, err := r.Client.Like.
		Create().
		SetOwnerID(like.UserID).
		SetLikedJokeID(like.JokeID).
		SetUserID(like.UserID).
		SetJokeID(like.JokeID).
		SetCreatedAt(like.CreatedAt).
		Save(ctx)
	return common.HandleErr(err)
}

func (r likeRepository) Delete(ctx context.Context, userID, jokeID string) error {
	_, err := r.Client.Like.
		Delete().
		Where(like.UserID(userID), like.JokeID(jokeID)).
		Exec(ctx)
	return common.HandleErr(err)
}
