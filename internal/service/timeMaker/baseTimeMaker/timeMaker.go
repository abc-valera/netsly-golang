package baseTimeMaker

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type baseTimeMaker struct {
}

func New() service.ITimeMaker {
	return baseTimeMaker{}
}

func (c baseTimeMaker) Now() time.Time {
	return time.Now()
}
