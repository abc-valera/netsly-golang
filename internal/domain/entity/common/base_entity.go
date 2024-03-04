package common

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/model/common"
	"github.com/google/uuid"
)

func NewBaseEntity() common.BaseEntity {
	return common.BaseEntity{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}
