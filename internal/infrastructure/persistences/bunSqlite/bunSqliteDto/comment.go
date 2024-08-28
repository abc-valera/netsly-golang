package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type Comment struct {
	ID        string    `bun:",pk,type:uuid"`
	Text      string    `bun:",notnull"`
	CreatedAt time.Time `bun:",notnull"`
	UpdatedAt time.Time `bun:",notnull"`
	DeletedAt time.Time `bun:",notnull"`

	UserID string `bun:",notnull"`
	JokeID string `bun:",notnull"`
}

func (c Comment) ToDomain() model.Comment {
	return model.Comment{
		ID:        c.ID,
		Text:      c.Text,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}

type Comments []Comment

func (c Comments) ToDomain() model.Comments {
	comments := make(model.Comments, 0, len(c))
	for _, comment := range c {
		comments = append(comments, comment.ToDomain())
	}
	return comments
}
