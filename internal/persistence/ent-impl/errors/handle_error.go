package errors

import (
	"strings"

	"github.com/abc-valera/netsly-api-golang/gen/ent"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model"
)

// HandleErr handles errors from db driver and converts them to domain errors
func HandleErr(err error) error {
	if err == nil {
		return nil
	}

	// Not found errors
	if ent.IsNotFound(err) {
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

		return coderr.NewMessage(coderr.CodeNotFound, "")
	}

	// Unique constraint errors
	if ent.IsConstraintError(err) {
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

		return coderr.NewMessage(coderr.CodeAlreadyExists, "")
	}

	return coderr.NewInternal(err)
}
