package newbasemodel

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
)

// Here functions now and newUUID are used to make the mocking possible.

func NewBaseModel(id string, createdAt time.Time) common.BaseModel {
	return common.BaseModel{
		ID:        id,
		CreatedAt: createdAt,
	}
}
