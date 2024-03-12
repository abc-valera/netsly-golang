package uuidMaker

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/google/uuid"
)

type uuidMaker struct {
}

func NewUUID() service.IUuidMaker {
	return uuidMaker{}
}

func (c uuidMaker) NewUUID() string {
	return uuid.NewString()
}
