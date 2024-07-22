package boilerSqlite

import (
	"database/sql"
	"embed"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/pressly/goose/v3"
)

//go:embed migration
var embedMigrations embed.FS

// New initializes the database connection and runs the migrations
func New(boilerSqlitePath string) (*sql.DB, error) {
	// Initialize the database connection
	db, err := sql.Open("sqlite", boilerSqlitePath)
	if err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	// Configure goosse and run the migrations
	goose.SetLogger(goose.NopLogger())
	goose.SetDialect(string(goose.DialectSQLite3))
	goose.SetBaseFS(embedMigrations)
	if err := goose.Up(db, "migration"); err != nil {
		return nil, coderr.NewInternalErr(err)
	}
	if err := goose.Version(db, "migration"); err != nil {
		return nil, coderr.NewInternalErr(err)
	}

	return db, nil
}
