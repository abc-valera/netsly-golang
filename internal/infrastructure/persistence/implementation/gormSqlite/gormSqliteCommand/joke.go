package gormSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteDto"
	"github.com/abc-valera/netsly-api-golang/internal/infrastructure/persistence/implementation/gormSqlite/gormSqliteErrutil"
	"gorm.io/gorm"
)

type joke struct {
	db *gorm.DB
}

func NewJoke(db *gorm.DB) command.IJoke {
	return &joke{
		db: db,
	}
}

func (c joke) Create(ctx context.Context, req command.JokeCreateRequest) (model.Joke, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	joke := gormSqliteDto.Joke{
		ID:          req.Joke.ID,
		Title:       req.Joke.Title,
		Text:        req.Joke.Text,
		Explanation: req.Joke.Explanation,
		CreatedAt:   req.Joke.CreatedAt,
		UpdatedAt:   req.Joke.UpdatedAt,
		DeletedAt:   req.Joke.DeletedAt,
		UserID:      req.UserID,
	}
	res := c.db.Create(&joke)
	return gormSqliteDto.NewDomainJoke(joke), gormSqliteErrutil.HandleCommandResult(res)
}

func (c joke) Update(ctx context.Context, id string, req command.JokeUpdateRequest) (model.Joke, error) {
	_, span := global.NewSpan(ctx)
	defer span.End()

	var joke gormSqliteDto.Joke
	queryRes := c.db.Where("id = ?", id).First(&joke)
	if err := gormSqliteErrutil.HandleQueryResult(queryRes); err != nil {
		return model.Joke{}, err
	}

	if req.Title != nil {
		joke.Title = *req.Title
	}
	if req.Text != nil {
		joke.Text = *req.Text
	}
	if req.Explanation != nil {
		joke.Explanation = *req.Explanation
	}

	updateRes := c.db.Save(&joke)
	return gormSqliteDto.NewDomainJoke(joke), gormSqliteErrutil.HandleCommandResult(updateRes)
}

func (c joke) Delete(ctx context.Context, id string) error {
	_, span := global.NewSpan(ctx)
	defer span.End()

	joke := gormSqliteDto.Joke{
		ID: id,
	}
	res := c.db.Delete(&joke)
	return gormSqliteErrutil.HandleCommandResult(res)
}
