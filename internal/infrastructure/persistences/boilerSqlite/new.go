package boilerSqlite

// Note, that we don't import sqlite driver here
// because it is already imported in the gormSqlite/new.go file.

import (
	"database/sql"
	"embed"
	"os"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/pressly/goose/v3"
)

//go:embed migration
var embedMigrations embed.FS

// New initializes the database connection and runs the migrations
func New(boilerSqliteFolderPath string) (*sql.DB, error) {
	// Create the folder
	if err := os.MkdirAll(boilerSqliteFolderPath, 0o755); err != nil {
		if !os.IsExist(err) {
			return nil, coderr.NewInternalErr(err)
		}
	}

	// Initialize the database connection
	db, err := sql.Open("sqlite", boilerSqliteFolderPath+"/sqlite.db")
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
