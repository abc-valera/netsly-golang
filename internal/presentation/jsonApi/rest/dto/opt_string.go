package dto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
)

func NewOptString(s string) ogen.OptString {
	if s == "" {
		return ogen.OptString{Set: false}
	}
	return ogen.NewOptString(s)
}

func NewPointerString(s ogen.OptString) *string {
	if s.Set {
		return &s.Value
	}
	return nil
}
