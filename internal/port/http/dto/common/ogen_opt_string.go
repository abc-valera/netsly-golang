package common

import (
	"github.com/abc-valera/flugo-api-golang/gen/ogen"
)

func NewOptString(s string) ogen.OptString {
	if s == "" {
		return ogen.OptString{Set: false}
	}
	return ogen.NewOptString(s)
}
