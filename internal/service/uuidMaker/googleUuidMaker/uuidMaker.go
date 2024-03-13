package googleUuidMaker

import (
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/google/uuid"
)

type googleUuidMaker struct {
}

func New() service.IUuidMaker {
	return googleUuidMaker{}
}

func (c googleUuidMaker) NewUUID() string {
	return uuid.NewString()
}
