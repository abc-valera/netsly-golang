package bunSqliteErrutil

import (
	"database/sql"

	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
)

func HandleCommandResult(res sql.Result, err error) error {
	if err != nil {
		// Handler errors here
		return coderr.NewInternalErr(err)
	}

	return nil
}
