package command

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command/commandGeneric"
)

type IFileContent commandGeneric.ICreateUpdateDelete[model.FileContent]
