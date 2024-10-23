package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/uptrace/bun"
)

type Joke struct {
	bun.BaseModel

	ID          string    `bun:"id,pk,type:uuid"`
	Title       string    `bun:",unique,notnull"`
	Text        string    `bun:",notnull"`
	Explanation string    `bun:",notnull"`
	CreatedAt   time.Time `bun:",notnull"`
	UpdatedAt   time.Time `bun:",notnull"`
	DeletedAt   time.Time `bun:",notnull"`

	UserID   string    `bun:",notnull"`
	Likes    []Like    `bun:"rel:has-many,join:id=joke_id"`
	Comments []Comment `bun:"rel:has-many,join:id=joke_id"`
}

func NewJoke(joke model.Joke) Joke {
	return Joke{
		ID:          joke.ID,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: joke.Explanation,
		CreatedAt:   joke.CreatedAt,
		UpdatedAt:   joke.UpdatedAt,
		DeletedAt:   joke.DeletedAt,

		UserID: joke.UserID,
	}
}

func (dto Joke) ToDomain() model.Joke {
	return model.Joke{
		ID:          dto.ID,
		Title:       dto.Title,
		Text:        dto.Text,
		Explanation: dto.Explanation,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
		DeletedAt:   dto.DeletedAt,

		UserID: dto.UserID,
	}
}
