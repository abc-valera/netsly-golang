package test

import (
	"github.com/abc-valera/netsly-api-golang/pkg/core/global"
	"github.com/abc-valera/netsly-api-golang/pkg/core/mode"
	"github.com/abc-valera/netsly-api-golang/pkg/service/logger/nopLogger"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// InitTestMain initializes DB connection
// and should be called in the beggining of every TestMain function in every test package
func InitTestMain() {
	// Init global
	global.InitLog(nopLogger.New())
	global.InitMode(mode.Production)
}
