package bunSqliteQuery

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteDto"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/persistences/bunSqlite/bunSqliteErrutil"
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
	return joke.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q joke) GetAllByUserID(ctx context.Context, userID string, selector selector.Selector) (model.Jokes, error) {
	jokes := bunSqliteDto.Jokes{}
	err := q.db.NewSelect().Model(&jokes).Where("user_id = ?", userID).Scan(ctx)
	return jokes.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}

func (q joke) SearchAllByTitle(ctx context.Context, keyword string, selector selector.Selector) (model.Jokes, error) {
	jokes := bunSqliteDto.Jokes{}
	err := q.db.NewSelect().Model(&jokes).Where("title LIKE ?", "%"+keyword+"%").Scan(ctx)
	return jokes.ToDomain(), bunSqliteErrutil.HandleQueryResult(err)
}
