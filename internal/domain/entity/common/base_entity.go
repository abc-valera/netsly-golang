package common

import (
	"time"

	"github.com/google/uuid"
)

type BaseEnity struct {
	ID        string
	CreatedAt time.Time
}

func NewBaseEnity() *BaseEnity {
	return &BaseEnity{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}
