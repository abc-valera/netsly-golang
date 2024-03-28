//go:build tools

package tools

import (
	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/pressly/goose/v3"
	_ "github.com/ziutek/mymysql/godrv"

	_ "github.com/volatiletech/sqlboiler/v4"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql"

	_ "github.com/vektra/mockery/v2"

	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
