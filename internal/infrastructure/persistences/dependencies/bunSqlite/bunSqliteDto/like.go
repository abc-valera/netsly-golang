package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/uptrace/bun"
)

type Like struct {
	bun.BaseModel

	CreatedAt time.Time `bun:",notnull"`
	DeletedAt time.Time `bun:",notnull"`

	UserID string `bun:",pk,notnull"`
	JokeID string `bun:",pk,notnull"`
}

func NewLike(like model.Like) Like {
	return Like{
		CreatedAt: like.CreatedAt,
		DeletedAt: like.DeletedAt,

		UserID: like.UserID,
		JokeID: like.JokeID,
	}
}

func (dto Like) ToDomain() model.Like {
	return model.Like{
		CreatedAt: dto.CreatedAt,
		DeletedAt: dto.DeletedAt,

		UserID: dto.UserID,
		JokeID: dto.JokeID,
	}
}
