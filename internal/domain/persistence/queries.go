package persistence

import "github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"

type Queries struct {
	User        query.IUser
	Joke        query.IJoke
	Like        query.ILike
	Comment     query.IComment
	Room        query.IRoom
	RoomMember  query.IRoomMember
	RoomMessage query.IRoomMessage
}
