package bunSqliteErrutil

import "github.com/abc-valera/netsly-golang/internal/core/coderr"

func HandleQueryResult(err error) error {
	if err != nil {
		// Handler errors here
		return coderr.NewInternalErr(err)
	}

	return nil
}
