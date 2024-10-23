package command

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
)

type IRoomMessage commandGeneric.ICreateUpdateDelete[model.RoomMessage]
