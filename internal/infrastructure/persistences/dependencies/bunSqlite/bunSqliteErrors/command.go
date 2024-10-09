package bunSqliteErrors

import (
	"database/sql"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

// TODO: add error handling here
func HandleCommandResult(res sql.Result, err error) error {
	_ = res

	if err != nil {
		// Handler errors here
		return coderr.NewInternalErr(err)
	}

	return nil
}
