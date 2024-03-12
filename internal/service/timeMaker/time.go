package timeMaker

import (
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
)

type timeMaker struct {
}

func NewTimeMaker() service.ITimeMaker {
	return timeMaker{}
}

func (c timeMaker) Now() time.Time {
	return time.Now()
}
