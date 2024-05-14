package restDto

import (
	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/core/optional"
)

func NewDomainOptionalString(opt ogen.OptString) optional.Optional[string] {
	if opt.Set {
		return optional.NewOptional(opt.Value)
	}
	return optional.NewEmptyOptional[string]()
}

func NewOgenOptionalString(opt optional.Optional[string]) ogen.OptString {
	if opt.IsPresent() {
		return ogen.OptString{Set: true, Value: opt.Value()}
	}
	return ogen.OptString{Set: false}
}
