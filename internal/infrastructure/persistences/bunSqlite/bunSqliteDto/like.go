package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type Like struct {
	CreatedAt time.Time `bun:",notnull"`
	DeletedAt time.Time `bun:",notnull"`

	UserID string `bun:",pk,notnull"`
	JokeID string `bun:",pk,notnull"`
}

func (l Like) ToDomain() model.Like {
	return model.Like{
		CreatedAt: l.CreatedAt,
		DeletedAt: l.DeletedAt,
	}
}

type Likes []Like

func (l Likes) ToDomain() model.Likes {
	likes := make(model.Likes, 0, len(l))
	for _, like := range l {
		likes = append(likes, like.ToDomain())
	}
	return likes
}
