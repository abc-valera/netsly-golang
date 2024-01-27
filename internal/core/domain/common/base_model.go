package common

import (
	"time"

	"github.com/abc-valera/flugo-api-golang/internal/core/model/common"
	"github.com/google/uuid"
)

func NewBaseModel() common.BaseModel {
	return common.BaseModel{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}
