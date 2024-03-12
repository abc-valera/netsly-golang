package sqlboilerImpl

import (
	"database/sql"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	_ "github.com/lib/pq"
)

func Init(postgresUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		return nil, coderr.NewInternalErr(err)
	}
	return db, nil
}
