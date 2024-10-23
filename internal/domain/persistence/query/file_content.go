package query

import (
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/queryGeneric"
)

type IFileContent queryGeneric.IGetOne[model.FileContent]
