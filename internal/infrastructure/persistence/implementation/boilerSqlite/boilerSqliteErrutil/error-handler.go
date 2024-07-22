package boilerSqliteErrutil

import (
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
)

// HandleErr handles errors from db driver and converts them to domain errors
func HandleErr(err error) error {
	if err == nil {
		return nil
	}

	// TODO: write error handling here
	return coderr.NewInternalErr(err)
}
