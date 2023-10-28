package common

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        string
	CreatedAt time.Time
}

func NewBaseEntity() BaseEntity {
	return BaseEntity{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}
