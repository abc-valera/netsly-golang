//go:build tools

package tools

import (
	// goose is used for managing database migrations.
	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/pressly/goose/v3"
	_ "github.com/ziutek/mymysql/godrv"

	// sqlboiler is used to generate ORM related code.
	_ "github.com/volatiletech/sqlboiler/v4"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql"

	// mockery is used for generating mocks for interfaces.
	_ "github.com/vektra/mockery/v2"

	// gqlgen is used to generate GraphQL related code.
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
