package bunSqliteCommand

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	bunSqlitErrutil "github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
	"github.com/uptrace/bun"
)

type joke struct {
	db bun.IDB
}

func NewJoke(db bun.IDB) command.IJoke {
	return &joke{
		db: db,
	}
}

func (c joke) Create(ctx context.Context, req command.JokeCreateRequest) (model.Joke, error) {
	joke := bunSqliteDto.Joke{
		ID:          req.Joke.ID,
		Title:       req.Joke.Title,
		Text:        req.Joke.Text,
		Explanation: req.Joke.Explanation,
		CreatedAt:   req.Joke.CreatedAt,
		UpdatedAt:   req.Joke.UpdatedAt,
		DeletedAt:   req.Joke.DeletedAt,

		UserID: req.UserID,
	}

	res, err := c.db.NewInsert().Model(&joke).Exec(ctx)
	return joke.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c joke) Update(ctx context.Context, id string, req command.JokeUpdateRequest) (model.Joke, error) {
	joke := bunSqliteDto.Joke{
		ID: id,
	}
	var columns []string

	if req.Title != nil {
		joke.Title = *req.Title
		columns = append(columns, "title")
	}
	if req.Text != nil {
		joke.Text = *req.Text
		columns = append(columns, "text")
	}
	if req.Explanation != nil {
		joke.Explanation = *req.Explanation
		columns = append(columns, "explanation")
	}

	if len(columns) == 0 {
		return model.Joke{}, nil
	}

	res, err := c.db.NewUpdate().Model(&joke).Column(columns...).WherePK().Exec(ctx)
	return joke.ToDomain(), bunSqlitErrutil.HandleCommandResult(res, err)
}

func (c joke) Delete(ctx context.Context, id string) error {
	joke := bunSqliteDto.Joke{
		ID: id,
	}
	res, err := c.db.NewDelete().Model(&joke).WherePK().Exec(ctx)
	return bunSqlitErrutil.HandleCommandResult(res, err)
}
