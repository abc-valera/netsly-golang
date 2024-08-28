package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type Joke struct {
	ID          string    `bun:",pk,type:uuid"`
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

func (j Joke) ToDomain() model.Joke {
	return model.Joke{
		ID:          j.ID,
		Title:       j.Title,
		Text:        j.Text,
		Explanation: j.Explanation,
		CreatedAt:   j.CreatedAt,
		UpdatedAt:   j.UpdatedAt,
		DeletedAt:   j.DeletedAt,
	}
}

type Jokes []Joke

func (j Jokes) ToDomain() model.Jokes {
	jokes := make(model.Jokes, 0, len(j))
	for _, joke := range j {
		jokes = append(jokes, joke.ToDomain())
	}
	return jokes
}
