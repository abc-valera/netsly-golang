package test

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

// InitTestMain initializes DB connection
// and should be called in the beggining of every TestMain function in every test package
func InitTestMain() {
	// Init global variables
	// global.InitMode(mode.Production)
	// global.InitLog(nopLogger.New())
}
