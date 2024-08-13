package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
)

type Joke struct {
	ID          string    `gorm:"primaryKey;not null"`
	Title       string    `gorm:"not null;index:unique_title_per_user,unique"`
	Text        string    `gorm:"not null"`
	Explanation string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	DeletedAt   time.Time `gorm:"not null"`

	UserID   string    `gorm:"not null;index:unique_title_per_user,unique"`
	Likes    []Like    `gorm:"foreignKey:JokeID;constraint:OnDelete:CASCADE"`
	Comments []Comment `gorm:"foreignKey:JokeID;constraint:OnDelete:CASCADE"`
}

func NewDomainJoke(joke Joke) model.Joke {
	return model.Joke{
		ID:          joke.ID,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: joke.Explanation,
		CreatedAt:   joke.CreatedAt,
		UpdatedAt:   joke.UpdatedAt,
		DeletedAt:   joke.DeletedAt,
	}
}

type Jokes []Joke

func NewDomainJokes(jokes Jokes) model.Jokes {
	var domainJokes model.Jokes
	for _, joke := range jokes {
		domainJokes = append(domainJokes, NewDomainJoke(joke))
	}
	return domainJokes
}
