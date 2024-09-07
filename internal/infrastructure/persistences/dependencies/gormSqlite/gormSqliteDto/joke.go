package gormSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
)

type Joke struct {
	ID          string    `gorm:"primaryKey;not null"`
	Title       string    `gorm:"not null;index:unique_title_per_user,unique"`
	Text        string    `gorm:"not null"`
	Explanation string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	DeletedAt   time.Time `gorm:"not null"`
	UserID      string    `gorm:"not null;index:unique_title_per_user,unique"`
	Likes       Likes     `gorm:"foreignKey:JokeID;constraint:OnDelete:CASCADE"`
	Comments    Comments  `gorm:"foreignKey:JokeID;constraint:OnDelete:CASCADE"`
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
		UserID:      joke.UserID,
	}
}

func NewJokeUpdate(joke Joke, req command.JokeUpdateRequest) Joke {
	joke.UpdatedAt = req.UpdatedAt

	if req.Title != nil {
		joke.Title = *req.Title
	}
	if req.Text != nil {
		joke.Text = *req.Text
	}
	if req.Explanation != nil {
		joke.Explanation = *req.Explanation
	}

	return joke
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
		UserID:      dto.UserID,
	}
}

type Jokes []Joke

func NewJokes(domainJokes model.Jokes) Jokes {
	var jokes Jokes
	for _, domainJoke := range domainJokes {
		jokes = append(jokes, NewJoke(domainJoke))
	}
	return jokes
}

func (dtos Jokes) ToDomain() model.Jokes {
	var domainJokes model.Jokes
	for _, dto := range dtos {
		domainJokes = append(domainJokes, dto.ToDomain())
	}
	return domainJokes
}
