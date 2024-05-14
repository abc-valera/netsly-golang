package errors

import (
	"strings"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/model"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
)

// HandleErr handles errors from db driver and converts them to domain errors
func HandleErr(err error) error {
	if err == nil {
		return nil
	}

	// TODO: write correct error clauses
	// Not found errors
	if strings.Contains(err.Error(), "no rows in result set") {
		// users table
		if strings.Contains(err.Error(), "user") {
			return model.ErrUserNotFound
		}

		// jokes table
		if strings.Contains(err.Error(), "joke") {
			return model.ErrJokeNotFound
		}

		// comments table
		if strings.Contains(err.Error(), "comment") {
			return model.ErrCommentNotFound
		}

		// likes table
		if strings.Contains(err.Error(), "like") {
			return model.ErrLikeNotFound
		}

		// chat rooms table
		if strings.Contains(err.Error(), "room") {
			return model.ErrRoomNotFound
		}

		// chat members table
		if strings.Contains(err.Error(), "chat_member") {
			return model.ErrRoomMemberNotFound
		}

		// chat msgs table
		if strings.Contains(err.Error(), "chat_msg") {
			return model.ErrRoomMessageNotFound
		}

		return coderr.NewCodeMessage(coderr.CodeNotFound, "")
	}

	// Unique constraint errors
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		// users table
		if strings.Contains(err.Error(), "users") {
			// username field
			if strings.Contains(err.Error(), "username") {
				return command.ErrUserWithUsernameAlreadyExists
			}
			// email field
			if strings.Contains(err.Error(), "email") {
				return command.ErrUserWithEmailAlreadyExists
			}
		}

		// joke-owner edge
		if strings.Contains(err.Error(), "joke") {
			return command.ErrJokeOwnerTitleAlreadyExists
		}

		// likes table
		if strings.Contains(err.Error(), "likes") {
			return command.ErrLikeAlreadyExists
		}

		// chat rooms table
		if strings.Contains(err.Error(), "rooms") {
			return command.ErrRoomNameAlreadyExists
		}

		// chat members table
		if strings.Contains(err.Error(), "chat_members") {
			return command.ErrRoomMemberAlreadyExists
		}

		return coderr.NewCodeMessage(coderr.CodeAlreadyExists, "")
	}

	return coderr.NewInternalErr(err)
}
