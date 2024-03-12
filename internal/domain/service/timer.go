package service

import "time"

type ITimeMaker interface {
	Now() time.Time
}
