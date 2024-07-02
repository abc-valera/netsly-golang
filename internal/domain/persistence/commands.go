package persistence

import "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/command"

type Commands struct {
	User        command.IUser
	Joke        command.IJoke
	Like        command.ILike
	Comment     command.IComment
	Room        command.IRoom
	RoomMember  command.IRoomMember
	RoomMessage command.IRoomMessage
	FileInfo    command.IFileInfo
}
