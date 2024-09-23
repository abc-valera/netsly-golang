package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/uptrace/bun"
)

type Joke struct {
	bun.BaseModel `bun:"table:jokes"`

	ID          string    `bun:"id,pk,type:uuid"`
	Title       string    `bun:",unique,notnull"`
	Text        string    `bun:",notnull"`
	Explanation string    `bun:",notnull"`
	CreatedAt   time.Time `bun:",notnull"`
	UpdatedAt   time.Time `bun:",notnull"`
	DeletedAt   time.Time `bun:",notnull"`

	UserID   string   `bun:",notnull"`
	Likes    Likes    `bun:"rel:has-many,join:id=joke_id"`
	Comments Comments `bun:"rel:has-many,join:id=joke_id"`
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

type Jokes []Joke

func NewJokes(jokes model.Jokes) Jokes {
	dtos := make(Jokes, 0, len(jokes))
	for _, joke := range jokes {
		dtos = append(dtos, NewJoke(joke))
	}
	return dtos
}

func (dtos Jokes) ToDomain() model.Jokes {
	jokes := make(model.Jokes, 0, len(dtos))
	for _, joke := range dtos {
		jokes = append(jokes, joke.ToDomain())
	}
	return jokes
}
