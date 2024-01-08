package domain

import (
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/command"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
)

type Domains struct {
	User        User
	Joke        Joke
	Like        Like
	Comment     Comment
	ChatRoom    ChatRoom
	ChatMember  ChatMember
	ChatMessage ChatMessage
}

func NewDomains(
	commands command.Commands,
	queries query.Queries,
	services service.Services,
) Domains {
	return Domains{
		User:        NewUser(commands.User, queries.User, services.PasswordMaker),
		Joke:        NewJoke(commands.Joke),
		Like:        NewLike(commands.Like),
		Comment:     NewComment(commands.Comment),
		ChatRoom:    NewChatRoom(commands.ChatRoom),
		ChatMember:  NewChatMember(commands.ChatMember),
		ChatMessage: NewChatMessage(commands.ChatMessage),
	}
}
