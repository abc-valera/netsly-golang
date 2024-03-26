package domain

import (
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/command"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
)

type Commands struct {
	User        command.IUser
	Joke        command.IJoke
	Like        command.ILike
	Comment     command.IComment
	Room        command.IRoom
	RoomMember  command.IRoomMember
	RoomMessage command.IRoomMessage
}

func NewCommands(
	user command.IUser,
	joke command.IJoke,
	like command.ILike,
	comment command.IComment,
	room command.IRoom,
	roomMember command.IRoomMember,
	roomMessage command.IRoomMessage,
) Commands {
	return Commands{
		User:        user,
		Joke:        joke,
		Like:        like,
		Comment:     comment,
		Room:        room,
		RoomMember:  roomMember,
		RoomMessage: roomMessage,
	}
}

type Queries struct {
	User        query.IUser
	Joke        query.IJoke
	Like        query.ILike
	Comment     query.IComment
	Room        query.IRoom
	RoomMember  query.IRoomMember
	RoomMessage query.IRoomMessage
}

func NewQueries(
	user query.IUser,
	joke query.IJoke,
	like query.ILike,
	comment query.IComment,
	room query.IRoom,
	roomMember query.IRoomMember,
	roomMessage query.IRoomMessage,
) Queries {
	return Queries{
		User:        user,
		Joke:        joke,
		Like:        like,
		Comment:     comment,
		Room:        room,
		RoomMember:  roomMember,
		RoomMessage: roomMessage,
	}
}
