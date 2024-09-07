package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/uptrace/bun"
)

type Comment struct {
	bun.BaseModel `bun:"table:comments"`

	ID        string    `bun:"id,pk,type:uuid"`
	Text      string    `bun:",notnull"`
	CreatedAt time.Time `bun:",notnull"`
	UpdatedAt time.Time `bun:",notnull"`
	DeletedAt time.Time `bun:",notnull"`

	UserID string `bun:",notnull"`
	JokeID string `bun:",notnull"`
}

func NewComment(comment model.Comment) Comment {
	return Comment{
		ID:        comment.ID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		DeletedAt: comment.DeletedAt,

		UserID: comment.UserID,
		JokeID: comment.JokeID,
	}
}

func NewCommentUpdate(ids model.Comment, req command.CommentUpdateRequest) (Comment, []string) {
	comment := Comment{
		ID: ids.ID,
	}
	var columns []string

	comment.UpdatedAt = req.UpdatedAt
	columns = append(columns, "updated_at")
	if req.Text != nil {
		comment.Text = *req.Text
		columns = append(columns, "text")
	}

	return comment, columns
}

func (dto Comment) ToDomain() model.Comment {
	return model.Comment{
		ID:        dto.ID,
		Text:      dto.Text,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,

		UserID: dto.UserID,
		JokeID: dto.JokeID,
	}
}

type Comments []Comment

func NewComments(comments model.Comments) Comments {
	dtos := make(Comments, 0, len(comments))
	for _, comment := range comments {
		dtos = append(dtos, NewComment(comment))
	}
	return dtos
}

func (dtos Comments) ToDomain() model.Comments {
	comments := make(model.Comments, 0, len(dtos))
	for _, comment := range dtos {
		comments = append(comments, comment.ToDomain())
	}
	return comments
}
