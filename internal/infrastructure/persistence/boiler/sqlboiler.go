package boiler

import (
	"database/sql"
	"embed"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migration/*.sql
var embedMigrations embed.FS

func Init(postgresUrl string) (*sql.DB, error) {
	// Initialize the database connection
	db, err := sql.Open("pgx", postgresUrl)
	if err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	// Configure goosse and run the migrations
	goose.SetLogger(goose.NopLogger())
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
