package impl

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/gen/ent/comment"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/dto"
	"github.com/abc-valera/flugo-api-golang/internal/adapter/persistence/ent/impl/common"
	"github.com/abc-valera/flugo-api-golang/internal/domain/entity"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository/spec"
)

type commentRepository struct {
	common.BaseRepository
}

func NewCommentRepository(client *ent.Client) repository.ICommentRepository {
	return &commentRepository{
		BaseRepository: common.NewBaseRepository(client),
	}
}

func (r commentRepository) Create(ctx context.Context, comment *entity.Comment) error {
	_, err := r.Client.Comment.
		Create().
		SetID(comment.ID).
		SetUserID(comment.UserID).
		SetJokeID(comment.JokeID).
		SetText(comment.Text).
		SetCreatedAt(comment.CreatedAt).
		Save(ctx)
	return common.HandleErr(err)
}

func (r commentRepository) GetByID(ctx context.Context, id string) (*entity.Comment, error) {
	entComment, err := r.Client.Comment.
		Query().
		Where(comment.ID(id)).
		Only(ctx)
	return dto.FromEntCommentToComment(entComment), common.HandleErr(err)
}

func (r commentRepository) GetByJokeID(ctx context.Context, jokeID string, spec spec.SelectParams) (entity.Comments, error) {
	entComments, err := r.specToQuery(spec).
		Where(comment.JokeID(jokeID)).
		All(ctx)
	return dto.FromEntCommentsToComments(entComments), common.HandleErr(err)
}

func (r commentRepository) Update(ctx context.Context, domainComment *entity.Comment) error {
	_, err := r.Client.Comment.
		Update().
		Where(comment.ID(domainComment.ID)).
		SetText(domainComment.Text).
		Save(ctx)
	return common.HandleErr(err)
}

func (r commentRepository) Delete(ctx context.Context, commentID string) error {
	_, err := r.Client.Comment.
		Delete().
		Where(comment.ID(commentID)).
		Exec(ctx)
	return common.HandleErr(err)
}

func (r commentRepository) specToQuery(spec spec.SelectParams) *ent.CommentQuery {
	return r.Client.Comment.
		Query().
		Order(ent.Desc(comment.FieldCreatedAt)).
		Limit(int(spec.Limit)).
		Offset(int(spec.Offset))
}
