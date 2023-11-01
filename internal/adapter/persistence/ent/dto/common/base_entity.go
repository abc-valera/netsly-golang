package common

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/domain/entity/common"
)

func FromEntToBaseEntity(id string, createdAt time.Time) common.BaseEntity {
	return common.BaseEntity{
		ID:        id,
		CreatedAt: createdAt,
	}
}
