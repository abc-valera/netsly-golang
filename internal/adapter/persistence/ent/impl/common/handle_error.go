package common

import (
	"strings"

	"github.com/abc-valera/flugo-api-golang/gen/ent"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
)

func HandleErr(err error) error {
	if err == nil {
		return nil
	}
	// Unique constraint errors
	if ent.IsConstraintError(err) {
		if strings.Contains(err.Error(), "users") {
			if strings.Contains(err.Error(), "username") {
				return repository.ErrUserWithUsernameAlreadyExists
			}
			if strings.Contains(err.Error(), "email") {
				return repository.ErrUserWithEmailAlreadyExists
			}
		}
		if strings.Contains(err.Error(), "jokes") {
			if strings.Contains(err.Error(), "title") {
				return repository.ErrJokeWithTitleAlreadyExists
			}
		}
		// all other tables...
		return codeerr.NewMsgErr(codeerr.CodeAlreadyExists, "")
	}
	// Not found errors
	if ent.IsNotFound(err) {
		if strings.Contains(err.Error(), "user") {
			return repository.ErrUserNotFound
		}
		if strings.Contains(err.Error(), "joke") {
			return repository.ErrJokeNotFound
		}
		// all other tables...
		return codeerr.NewMsgErr(codeerr.CodeNotFound, "")
	}
	return codeerr.NewInternal("ent.handleErr", err)
}
