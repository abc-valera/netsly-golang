package coderr

import (
	"bytes"
	"errors"
	"fmt"
)

type internal struct {
	Caller string // Caller provides additional context about error's location
	Err    error  // Err is a nested error
}

func NewInternalErr(err error) error {
	return &internal{
		Caller: caller(2),
		Err:    err,
	}
}

func NewInternalString(err string) error {
	return &internal{
		Caller: caller(2),
		Err:    errors.New(err),
	}
}

func (e *internal) Error() string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "%s ", e.Caller)

	fmt.Fprintf(&buf, "<%s> ", CodeInternal)
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	}

	return buf.String()
}
