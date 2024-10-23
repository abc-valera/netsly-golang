package command

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
)

type IFileInfo commandGeneric.ICreateUpdateDelete[model.FileInfo]

type IFileInfoJoke commandGeneric.ICreate[model.FileInfoJoke]

type IFileInfoRoom commandGeneric.ICreate[model.FileInfoRoom]
