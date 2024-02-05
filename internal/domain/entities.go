package domain

import "github.com/abc-valera/netsly-api-golang/internal/domain/entity"

type Entities struct {
	User        entity.User
	Joke        entity.Joke
	Like        entity.Like
	Comment     entity.Comment
	ChatRoom    entity.ChatRoom
	ChatMember  entity.ChatMember
	ChatMessage entity.ChatMessage
}

func NewEntities(
	commands Commands,
	queries Queries,
	services Services,
) Entities {
	return Entities{
		User:        entity.NewUser(commands.User, queries.User, services.PasswordMaker),
		Joke:        entity.NewJoke(commands.Joke),
		Like:        entity.NewLike(commands.Like),
		Comment:     entity.NewComment(commands.Comment),
		ChatRoom:    entity.NewChatRoom(commands.ChatRoom),
		ChatMember:  entity.NewChatMember(commands.ChatMember),
		ChatMessage: entity.NewChatMessage(commands.ChatMessage),
	}
}
