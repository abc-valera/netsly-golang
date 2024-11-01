package command

type Commands struct {
	User        IUser
	Joke        IJoke
	Like        ILike
	Comment     IComment
	Room        IRoom
	RoomMember  IRoomMember
	RoomMessage IRoomMessage
	FileInfo    IFileInfo
	FileContent IFileContent
}

// TODO: maybe define errors that can be returned by the commands here
