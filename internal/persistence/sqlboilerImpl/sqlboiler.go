package sqlboilerImpl

import (
	"database/sql"
	"embed"

	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migration/*.sql
var embedMigrations embed.FS

func Init(postgresUrl string) (*sql.DB, error) {
	// Initialize the database connection
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	// Run the migrations
	// TODO: add logger for goose
	goose.SetDialect(string(goose.DialectPostgres))
	goose.SetBaseFS(embedMigrations)
	if err := goose.Up(db, "migration"); err != nil {
		return nil, coderr.NewInternalErr(err)
	}
	if err := goose.Version(db, "migration"); err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	return db, nil
}
