package common

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/core/persistence/model/common"
	"github.com/google/uuid"
)

func NewBaseModel() common.BaseModel {
	return common.BaseModel{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}
