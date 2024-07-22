package persistence

import (
	"fmt"
	"reflect"

	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
)

type Queries struct {
	User        query.IUser
	Joke        query.IJoke
	Like        query.ILike
	Comment     query.IComment
	Room        query.IRoom
	RoomMember  query.IRoomMember
	RoomMessage query.IRoomMessage
	FileInfo    query.IFileInfo
	FileContent query.IFileContent
}

func ValidateQueries(c Queries) error {
	defer func() {
		if r := recover(); r != nil {
			coderr.NewInternalString(fmt.Sprintf("Recovered: %v", r))
		}
	}()

	reflectVal := reflect.ValueOf(c)

	for i := range reflectVal.NumField() {
		field := reflectVal.Field(i)

		if field.Interface() == nil {
			return coderr.NewInternalString(reflectVal.Type().Field(0).Name + "Query is nil")
		} else {
			continue
		}
	}

	return nil
}
