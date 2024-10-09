package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/dependencies/bunSqlite/bunSqliteErrors"
	"github.com/uptrace/bun"
)

type joke struct {
	db bun.IDB
}

func NewJoke(db bun.IDB) query.IJoke {
	return &joke{
		db: db,
	}
}

func (q joke) GetByID(ctx context.Context, id string) (model.Joke, error) {
	joke := bunSqliteDto.Joke{}
	err := q.db.NewSelect().Model(&joke).Where("id = ?", id).Scan(ctx)
	return joke.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}

func (q joke) GetAllByUserID(ctx context.Context, userID string, s selector.Selector) (model.Jokes, error) {
	jokes := bunSqliteDto.Jokes{}
	err := q.db.NewSelect().Model(&jokes).Where("user_id = ?", userID).Scan(ctx)
	return jokes.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}

func (q joke) SearchAllByTitle(ctx context.Context, keyword string, s selector.Selector) (model.Jokes, error) {
	jokes := bunSqliteDto.Jokes{}
	err := q.db.NewSelect().Model(&jokes).Where("title LIKE ?", "%"+keyword+"%").Scan(ctx)
	return jokes.ToDomain(), bunSqliteErrors.HandleQueryResult(err)
}
