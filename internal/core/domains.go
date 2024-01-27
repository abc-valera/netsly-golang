package core

import "github.com/abc-valera/flugo-api-golang/internal/core/domain"

type Domains struct {
	User        domain.User
	Joke        domain.Joke
	Like        domain.Like
	Comment     domain.Comment
	ChatRoom    domain.ChatRoom
	ChatMember  domain.ChatMember
	ChatMessage domain.ChatMessage
}

func NewDomains(
	commands Commands,
	queries Queries,
	services Services,
) Domains {
	return Domains{
		User:        domain.NewUser(commands.User, queries.User, services.PasswordMaker),
		Joke:        domain.NewJoke(commands.Joke),
		Like:        domain.NewLike(commands.Like),
		Comment:     domain.NewComment(commands.Comment),
		ChatRoom:    domain.NewChatRoom(commands.ChatRoom),
		ChatMember:  domain.NewChatMember(commands.ChatMember),
		ChatMessage: domain.NewChatMessage(commands.ChatMessage),
	}
}
