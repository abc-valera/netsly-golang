package restDto

import (
	"github.com/abc-valera/netsly-golang/gen/ogen"
)

func NewDomainOptionalString(opt ogen.OptString) *string {
	if opt.Set {
		return &opt.Value
	}
	return nil
}

func NewOptionalString(opt *string) ogen.OptString {
	if opt != nil {
		return ogen.OptString{Set: true, Value: *opt}
	}
	return ogen.OptString{Set: false}
}
