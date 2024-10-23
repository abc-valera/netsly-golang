package restDto

import (
	"github.com/abc-valera/netsly-golang/gen/ogen"
)

func NewInt(opt ogen.OptInt) *int {
	if opt.Set {
		return &opt.Value
	}
	return nil
}

func NewOptionalInt(opt int) ogen.OptInt {
	return ogen.OptInt{Set: true, Value: opt}
}

func NewString(opt ogen.OptString) *string {
	if opt.Set {
		return &opt.Value
	}
	return nil
}

func NewOptionalString(opt string) ogen.OptString {
	return ogen.OptString{Set: true, Value: opt}
}
