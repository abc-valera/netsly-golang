package common

import (
	"time"
)

type BaseEntity struct {
	ID        string
	CreatedAt time.Time
}
