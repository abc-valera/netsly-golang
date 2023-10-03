package impl

import (
	"strings"

	"github.com/abc-valera/flugo-api-golang/internal/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/domain/repository"
	"github.com/abc-valera/flugo-api-golang/tools/ent"
)

func handleErr(err error) error {
	if err == nil {
		return nil
	}
	if ent.IsConstraintError(err) {
		if strings.Contains(err.Error(), "users") {
			if strings.Contains(err.Error(), "username") {
				return repository.ErrUserWithUsernameAlreadyExists
			}
			if strings.Contains(err.Error(), "email") {
				return repository.ErrUserWithEmailAlreadyExists
			}
		}
		// all other tables...
		return codeerr.NewErrWithMsg(codeerr.CodeAlreadyExists, "")
	}
	if ent.IsNotFound(err) {
		if strings.Contains(err.Error(), "user") {
			return repository.ErrUserNotFound
		}
		// all other tables...
		return codeerr.NewErrWithMsg(codeerr.CodeNotFound, "")
	}
	return codeerr.NewInternal("ent.handleErr", err)
}
