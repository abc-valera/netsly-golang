//go:build tools

// Here all the binaries that are used in the project are improted.
// This is a trick to make sure that the binaries are installed with a correct version.
// Thanks to the `//go:build` tag this file is not included in the final binary.

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
)
