package command

import (
	"errors"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/codeerr"
)

type Commands struct {
	User        IUser
	Joke        IJoke
	Like        ILike
	Comment     IComment
	ChatRoom    IChatRoom
	ChatMember  IChatMember
	ChatMessage IChatMessage
}

func NewCommands(
	user IUser,
	joke IJoke,
	like ILike,
	comment IComment,
	chatRoom IChatRoom,
	chatMember IChatMember,
	chatMessage IChatMessage,
) (Commands, error) {
	if user == nil {
		return Commands{}, codeerr.NewInternal(errors.New("user command is nil"))
	}
	if joke == nil {
		return Commands{}, codeerr.NewInternal(errors.New("joke command is nil"))
	}
	if like == nil {
		return Commands{}, codeerr.NewInternal(errors.New("like command is nil"))
	}
	if comment == nil {
		return Commands{}, codeerr.NewInternal(errors.New("comment command is nil"))
	}
	if chatRoom == nil {
		return Commands{}, codeerr.NewInternal(errors.New("chatRoom ommand is nil"))
	}
	if chatMember == nil {
		return Commands{}, codeerr.NewInternal(errors.New("chatMember command is nil"))
	}
	if chatMessage == nil {
		return Commands{}, codeerr.NewInternal(errors.New("chatMessage command is nil"))
	}
	return Commands{
		User:        user,
		Joke:        joke,
		Like:        like,
		Comment:     comment,
		ChatRoom:    chatRoom,
		ChatMember:  chatMember,
		ChatMessage: chatMessage,
	}, nil
}
